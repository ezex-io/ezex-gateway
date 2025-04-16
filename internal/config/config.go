package config

import (
	"github.com/ezex-io/ezex-gateway/api/graphql"
	"github.com/ezex-io/ezex-gateway/internal/adapter/grpc/notification"
	"github.com/ezex-io/ezex-gateway/internal/adapter/redis"
	"github.com/ezex-io/ezex-gateway/internal/interactor/auth"
	"github.com/ezex-io/ezex-gateway/internal/utils"
	"github.com/joho/godotenv"
)

type Config struct {
	Debug                     bool
	GraphqlConfig             *graphql.Config
	AuthInteractorConfig      *auth.Config
	NotificationAdapterConfig *notification.Config
	RedisAdapterConfig        *redis.Config
}

func LoadConfig(envFile string) (*Config, error) {
	// Load variables from .env file
	if err := godotenv.Load(envFile); err != nil {
		return nil, err
	}

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
		Debug:                     utils.GetEnvBoolOrDefault("DEBUG", false),
		GraphqlConfig:             graphqlConfig,
		AuthInteractorConfig:      authConfig,
		NotificationAdapterConfig: notificationConfig,
		RedisAdapterConfig:        redisConfig,
	}

	return config, nil
}

func (*Config) BasicCheck() error {
	// Add any necessary validation here
	return nil
}
