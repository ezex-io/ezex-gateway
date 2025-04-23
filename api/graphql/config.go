package graphql

import (
	"github.com/ezex-io/gopkg/env"
)

type Config struct {
	Address    string
	Playground bool
	QueryPath  string
	CORS       Cors
}

type Cors struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	AllowCredentials bool
}

func LoadFromEnv() (*Config, error) {
	config := &Config{
		Address:    env.GetEnv[string]("EZEX_GATEWAY_GRAPHQL_ADDRESS", env.WithDefault("0.0.0.0:8080")),
		QueryPath:  env.GetEnv[string]("EZEX_GATEWAY_GRAPHQL_QUERY_PATH"),
		CORS: Cors{
			AllowedOrigins:   env.GetEnv[[]string]("EZEX_GATEWAY_GRAPHQL_CORS_ALLOWED_ORIGINS"),
			AllowedMethods:   env.GetEnv[[]string]("EZEX_GATEWAY_GRAPHQL_CORS_ALLOWED_METHODS"),
			AllowedHeaders:   env.GetEnv[[]string]("EZEX_GATEWAY_GRAPHQL_CORS_ALLOWED_HEADERS"),
			AllowCredentials: env.GetEnv[bool]("EZEX_GATEWAY_GRAPHQL_CORS_ALLOW_CREDENTIALS", env.WithDefault("false")),
		},
	}

	return config, nil
}

func (*Config) BasicCheck() error {
	return nil
}
