package notification

import (
	"github.com/ezex-io/ezex-gateway/internal/utils"
)

type Config struct {
	Address string
	Port    int
}

func LoadFromEnv() (*Config, error) {
	config := &Config{
		Address: utils.GetEnvOrDefault("EZEX_GRPC_NOTIFICATION_ADDRESS", "0.0.0.0"),
		Port:    utils.GetEnvIntOrDefault("EZEX_GRPC_NOTIFICATION_PORT", 50051),
	}

	return config, nil
}

func (*Config) BasicCheck() error {
	return nil
}
