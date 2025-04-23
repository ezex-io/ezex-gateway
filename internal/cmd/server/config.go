package main

import (
	"github.com/ezex-io/ezex-gateway/api/graphql"
	"github.com/ezex-io/ezex-gateway/internal/adapter/grpc/notification"
	"github.com/ezex-io/ezex-gateway/internal/adapter/redis"
	"github.com/ezex-io/ezex-gateway/internal/interactor/auth"
	"github.com/ezex-io/gopkg/env"
)

type Config struct {
	Debug                     bool
	GraphqlConfig             *graphql.Config
	AuthInteractorConfig      *auth.Config
	NotificationAdapterConfig *notification.Config
	RedisAdapterConfig        *redis.Config
}

func makeConfig() (*Config, error) {
	graphqlConfig, err := graphql.LoadFromEnv()
	if err != nil {
		return nil, err
	}

	authConfig, err := auth.LoadFromEnv()
	if err != nil {
		return nil, err
	}

	notificationConfig, err := notification.LoadFromEnv()
	if err != nil {
		return nil, err
	}

	redisConfig, err := redis.LoadFromEnv()
	if err != nil {
		return nil, err
	}

	// Initialize config with environment variables
	config := &Config{
		Debug:                     env.GetEnv[bool]("DEBUG", env.WithDefault("false")),
		GraphqlConfig:             graphqlConfig,
		AuthInteractorConfig:      authConfig,
		NotificationAdapterConfig: notificationConfig,
		RedisAdapterConfig:        redisConfig,
	}

	config.GraphqlConfig.Playground = config.Debug

	return config, nil
}

func (*Config) BasicCheck() error {
	// Add any necessary validation here
	return nil
}
