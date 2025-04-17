package redis

import (
	"time"

	"github.com/ezex-io/ezex-gateway/internal/utils"
)

type Config struct {
	Host         string
	Port         int
	DB           int
	Password     string
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PoolSize     int
	Protocol     int
}

func LoadFromEnv() (*Config, error) {
	config := &Config{
		Host:         utils.GetEnvOrDefault("EZEX_REDIS_HOST", "localhost"),
		Port:         utils.GetEnvIntOrDefault("EZEX_REDIS_PORT", 6379),
		DB:           utils.GetEnvIntOrDefault("EZEX_REDIS_DB", 0),
		Password:     utils.GetEnvOrDefault("EZEX_REDIS_PASSWORD", ""),
		DialTimeout:  utils.GetEnvDurationOrDefault("EZEX_REDIS_DIAL_TIMEOUT", 5*time.Second),
		ReadTimeout:  utils.GetEnvDurationOrDefault("EZEX_REDIS_READ_TIMEOUT", 5*time.Second),
		WriteTimeout: utils.GetEnvDurationOrDefault("EZEX_REDIS_WRITE_TIMEOUT", 5*time.Second),
		PoolSize:     utils.GetEnvIntOrDefault("EZEX_REDIS_POOL_SIZE", 10),
		Protocol:     utils.GetEnvIntOrDefault("EZEX_REDIS_PROTOCOL", 3),
	}

	return config, nil
}

func (*Config) BasicCheck() error {
	return nil
}
