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

var DefaultConfig = &Config{
	Address:    "0.0.0.0",
	Port:       8080,
	Playground: true,
	QueryPath:  "",
	CORS: Cors{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	},
}

func LoadFromEnv() (*Config, error) {
	config := &Config{
		Address:    utils.GetEnvOrDefault("GRAPHQL_ADDRESS", DefaultConfig.Address),
		Port:       utils.GetEnvIntOrDefault("GRAPHQL_PORT", DefaultConfig.Port),
		Playground: utils.GetEnvBoolOrDefault("GRAPHQL_PLAYGROUND", DefaultConfig.Playground),
		QueryPath:  utils.GetEnvOrDefault("GRAPHQL_QUERY_PATH", DefaultConfig.QueryPath),
		CORS: Cors{
			AllowedOrigins:   utils.GetEnvSliceOrDefault("GRAPHQL_CORS_ALLOWED_ORIGINS", DefaultConfig.CORS.AllowedOrigins),
			AllowedMethods:   utils.GetEnvSliceOrDefault("GRAPHQL_CORS_ALLOWED_METHODS", DefaultConfig.CORS.AllowedMethods),
			AllowedHeaders:   utils.GetEnvSliceOrDefault("GRAPHQL_CORS_ALLOWED_HEADERS", DefaultConfig.CORS.AllowedHeaders),
			AllowCredentials: utils.GetEnvBoolOrDefault("GRAPHQL_CORS_ALLOW_CREDENTIALS", DefaultConfig.CORS.AllowCredentials),
		},
	}

	return config, nil
}

func (*Config) BasicCheck() error {
	return nil
}
