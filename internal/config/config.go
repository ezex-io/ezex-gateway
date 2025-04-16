package config

import (
	"time"

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

	// Initialize config with environment variables
	config := &Config{
		Debug: utils.GetEnvBoolOrDefault("DEBUG", false),
		GraphqlConfig: &graphql.Config{
			Address:    utils.GetEnvOrDefault("GRAPHQL_ADDRESS", "0.0.0.0"),
			Port:       utils.GetEnvIntOrDefault("GRAPHQL_PORT", 8080),
			Playground: utils.GetEnvBoolOrDefault("GRAPHQL_PLAYGROUND", true),
			QueryPath:  utils.GetEnvOrDefault("GRAPHQL_QUERY_PATH", ""),
			CORS: graphql.Cors{
				AllowedOrigins:   utils.GetEnvSliceOrDefault("GRAPHQL_CORS_ALLOWED_ORIGINS", []string{"*"}),
				AllowedMethods:   utils.GetEnvSliceOrDefault("GRAPHQL_CORS_ALLOWED_METHODS", []string{"GET", "POST", "PUT", "PATCH", "DELETE"}),
				AllowedHeaders:   utils.GetEnvSliceOrDefault("GRAPHQL_CORS_ALLOWED_HEADERS", []string{"*"}),
				AllowCredentials: utils.GetEnvBoolOrDefault("GRAPHQL_CORS_ALLOW_CREDENTIALS", true),
			},
		},
		AuthInteractorConfig: &auth.Config{
			ConfirmationCodeTTL:      utils.GetEnvDurationOrDefault("AUTH_CONFIRMATION_CODE_TTL", 5*time.Minute),
			ConfirmationTemplateName: utils.GetEnvOrDefault("AUTH_CONFIRMATION_TEMPLATE", "confirmation_letter"),
			ConfirmationCodeSubject:  utils.GetEnvOrDefault("AUTH_CONFIRMATION_SUBJECT", "ezeX Confirmation Code: %s"),
		},
		NotificationAdapterConfig: &notification.Config{
			Address: utils.GetEnvOrDefault("NOTIFICATION_ADDRESS", "0.0.0.0"),
			Port:    utils.GetEnvIntOrDefault("NOTIFICATION_PORT", 50051),
		},
		RedisAdapterConfig: &redis.Config{
			Host:         utils.GetEnvOrDefault("REDIS_HOST", "localhost"),
			Port:         utils.GetEnvIntOrDefault("REDIS_PORT", 6379),
			DB:           utils.GetEnvIntOrDefault("REDIS_DB", 0),
			Password:     utils.GetEnvOrDefault("REDIS_PASSWORD", ""),
			DialTimeout:  utils.GetEnvDurationOrDefault("REDIS_DIAL_TIMEOUT", 5*time.Second),
			ReadTimeout:  utils.GetEnvDurationOrDefault("REDIS_READ_TIMEOUT", 5*time.Second),
			WriteTimeout: utils.GetEnvDurationOrDefault("REDIS_WRITE_TIMEOUT", 5*time.Second),
			PoolSize:     utils.GetEnvIntOrDefault("REDIS_POOL_SIZE", 10),
			Protocol:     utils.GetEnvIntOrDefault("REDIS_PROTOCOL", 3),
		},
	}

	return config, nil
}

func (c *Config) BasicCheck() error {
	// Add any necessary validation here
	return nil
}
