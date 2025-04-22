package graphql

import (
	"github.com/ezex-io/ezex-gateway/internal/utils"
)

type Config struct {
	Address    string
	Port       int
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
		Address:    utils.GetEnvOrDefault("EZEX_GATEWAY_GRAPHQL_ADDRESS", "0.0.0.0"),
		Port:       utils.GetEnvIntOrDefault("EZEX_GATEWAY_GRAPHQL_PORT", 8080),
		Playground: utils.GetEnvBoolOrDefault("EZEX_GATEWAY_GRAPHQL_PLAYGROUND", true),
		QueryPath:  utils.GetEnvOrDefault("EZEX_GATEWAY_GRAPHQL_QUERY_PATH", ""),
		CORS: Cors{
			AllowedOrigins: utils.GetEnvSliceOrDefault(
				"EZEX_GATEWAY_GRAPHQL_CORS_ALLOWED_ORIGINS", []string{"*"},
			),
			AllowedMethods: utils.GetEnvSliceOrDefault(
				"EZEX_GATEWAY_GRAPHQL_CORS_ALLOWED_METHODS",
				[]string{"GET", "POST", "PUT", "PATCH", "DELETE"},
			),
			AllowedHeaders: utils.GetEnvSliceOrDefault(
				"EZEX_GATEWAY_GRAPHQL_CORS_ALLOWED_HEADERS", []string{"*"},
			),
			AllowCredentials: utils.GetEnvBoolOrDefault(
				"EZEX_GATEWAY_GRAPHQL_CORS_ALLOW_CREDENTIALS", true,
			),
		},
	}

	return config, nil
}

func (*Config) BasicCheck() error {
	return nil
}
