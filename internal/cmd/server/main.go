package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ezex-io/ezex-gateway/internal/adapter/ezex_notification"
	"github.com/ezex-io/ezex-gateway/internal/adapter/ezex_users"
	"github.com/ezex-io/ezex-gateway/internal/adapter/firebase"
	"github.com/ezex-io/ezex-gateway/internal/adapter/graphql"
	"github.com/ezex-io/ezex-gateway/internal/adapter/graphql/resolver"
	"github.com/ezex-io/ezex-gateway/internal/adapter/redis"
	"github.com/ezex-io/ezex-gateway/internal/interactor/auth"
	"github.com/ezex-io/gopkg/env"
	"github.com/ezex-io/gopkg/logger"
	mdl "github.com/ezex-io/gopkg/middleware/http-mdl"
)

func main() {
	envFile := flag.String("env", ".env", "Path to environment file")
	flag.Parse()

	logging := logger.NewSlog(nil)

	// TODO: move me into makeConfig
	if err := env.LoadEnvsFromFile(*envFile); err != nil {
		logging.Debug("Failed to load env file '%s': %v. Continuing with system environment...", *envFile, err)
	}

	cfg := makeConfig()
	logging.Info("successfully loaded config", "debug", cfg.Debug)

	if cfg.Debug {
		logging = logger.NewSlog(logger.WithTextHandler(os.Stdout, slog.LevelDebug))
	}

	redisPort, err := redis.New(cfg.Redis)
	if err != nil {
		logging.Fatal(err.Error())
	}
	logging.Info("initialized redis adapter")

	firebasePort, err := firebase.New(context.Background(), cfg.Firebase)
	if err != nil {
		logging.Fatal(err.Error())
	}

	notificationPort, err := ezex_notification.New(cfg.Notification)
	if err != nil {
		logging.Fatal(err.Error())
	}

	userPort, err := ezex_users.New(cfg.User)
	if err != nil {
		logging.Fatal(err.Error())
	}

	logging.Info("initialized notification service adapter")

	authInteractor := auth.NewAuth(cfg.AuthInteractor, logging, notificationPort, redisPort, firebasePort, userPort)

	resolver := resolver.NewResolver(authInteractor)

	gql := graphql.New(cfg.Graphql, resolver, logging, mdl.Recover())

	gql.Start()
	logging.Info("graphql server started", "addr", cfg.Graphql.Address)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case <-interrupt:
		err = gql.Stop(context.Background())
		if err != nil {
			logging.Fatal(err.Error())
		}

		err = redisPort.Close()
		if err != nil {
			logging.Fatal(err.Error())
		}

		err = notificationPort.Close()
		if err != nil {
			logging.Fatal(err.Error())
		}

		logging.Warn("service interrupted")
	case err := <-gql.Notify():
		logging.Error("graphql server got error", "err", err)
	}
}
