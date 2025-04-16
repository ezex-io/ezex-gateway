package notification

import (
	"github.com/ezex-io/ezex-gateway/internal/utils"
)

type Config struct {
	Address string
	Port    int
}

var DefaultConfig = &Config{
	Address: "0.0.0.0",
	Port:    50051,
}

func LoadFromEnv() (*Config, error) {
	config := &Config{
		Address: utils.GetEnvOrDefault("GRPC_NOTIFICATION_ADDRESS", DefaultConfig.Address),
		Port:    utils.GetEnvIntOrDefault("GRPC_NOTIFICATION_PORT", DefaultConfig.Port),
	}

	return config, nil
}

func (*Config) BasicCheck() error {
	return nil
}
