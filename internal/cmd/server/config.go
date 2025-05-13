package main

import (
	"github.com/ezex-io/ezex-gateway/internal/adapter/ezex_notification"
	"github.com/ezex-io/ezex-gateway/internal/adapter/ezex_users"
	"github.com/ezex-io/ezex-gateway/internal/adapter/firebase"
	"github.com/ezex-io/ezex-gateway/internal/adapter/graphql"
	"github.com/ezex-io/ezex-gateway/internal/adapter/redis"
	"github.com/ezex-io/ezex-gateway/internal/interactor/auth"
	"github.com/ezex-io/gopkg/env"
)

type Config struct {
	Debug          bool
	Graphql        *graphql.Config
	AuthInteractor *auth.Config
	Notification   *ezex_notification.Config
	User           *ezex_users.Config
	Redis          *redis.Config
	Firebase       *firebase.Config
}

func makeConfig() *Config {
	config := &Config{
		Debug:          env.GetEnv[bool]("EZEX_GATEWAY_DEBUG", env.WithDefault("false")),
		Graphql:        graphql.LoadFromEnv(),
		AuthInteractor: auth.LoadFromEnv(),
		Notification:   ezex_notification.LoadFromEnv(),
		Redis:          redis.LoadFromEnv(),
		Firebase:       firebase.LoadFromEnv(),
		User:           ezex_users.LoadFromEnv(),
	}

	config.Graphql.Playground = config.Debug

	return config
}

func (*Config) BasicCheck() error {
	// Add any necessary validation here
	return nil
}
