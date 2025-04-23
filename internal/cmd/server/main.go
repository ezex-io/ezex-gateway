package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ezex-io/ezex-gateway/api/graphql"
	"github.com/ezex-io/ezex-gateway/api/graphql/resolver"
	"github.com/ezex-io/ezex-gateway/internal/adapter/grpc/notification"
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

	cfg, err := makeConfig()
	if err != nil {
		logging.Fatal(err.Error())
	}
	logging.Info("successfully loaded config")

	if cfg.Debug {
		logging = logger.NewSlog(logger.WithTextHandler(os.Stdout, slog.LevelDebug))
	}

	redisPort, err := redis.New(cfg.RedisAdapterConfig)
	if err != nil {
		logging.Error(err.Error())
		os.Exit(1)
	}
	logging.Info("initialized redis adapter")

	notificationPort, err := notification.New(cfg.NotificationAdapterConfig)
	if err != nil {
		logging.Error(err.Error())
		os.Exit(1)
	}

	logging.Info("initialized notification service adapter")

	authInteractor := auth.NewAuth(cfg.AuthInteractorConfig, logging, notificationPort, redisPort)

	resolve := resolver.NewResolver(authInteractor)

	gql := graphql.New(cfg.GraphqlConfig, resolve, logging, mdl.Recover())

	gql.Start()
	logging.Info("graphql server started", "addr",
		fmt.Sprintf("%s:%d", cfg.GraphqlConfig.Address, cfg.GraphqlConfig.Port))

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
