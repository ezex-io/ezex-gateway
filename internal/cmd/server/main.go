package main

import (
	"context"
	"flag"
	"log/slog"
	"os"

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
	"github.com/ezex-io/gopkg/utils"
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

	ctx := context.Background()
	authenticatorPort, err := firebase.New(ctx, cfg.Firebase)
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

	authInteractor := auth.NewAuth(cfg.AuthInteractor, logging, notificationPort, redisPort, authenticatorPort, userPort)

	resolver := resolver.NewResolver(authInteractor)

	gql := graphql.New(cfg.Graphql, logging, resolver, mdl.Recover())

	gql.Start()

	utils.TrapSignal(func() {
		logging.Info("Exiting...")

		gql.Stop(ctx)

		redisPort.Close()
		notificationPort.Close()
		userPort.Close()
	})

	// run forever
	select {}
}
