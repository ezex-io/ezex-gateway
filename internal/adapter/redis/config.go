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

var DefaultConfig = &Config{
	Host:         "localhost",
	Port:         6379,
	DB:           0,
	Password:     "",
	DialTimeout:  5 * time.Second,
	ReadTimeout:  5 * time.Second,
	WriteTimeout: 5 * time.Second,
	PoolSize:     10,
	Protocol:     3,
}

func LoadFromEnv() (*Config, error) {
	config := &Config{
		Host:         utils.GetEnvOrDefault("REDIS_HOST", DefaultConfig.Host),
		Port:         utils.GetEnvIntOrDefault("REDIS_PORT", DefaultConfig.Port),
		DB:           utils.GetEnvIntOrDefault("REDIS_DB", DefaultConfig.DB),
		Password:     utils.GetEnvOrDefault("REDIS_PASSWORD", DefaultConfig.Password),
		DialTimeout:  utils.GetEnvDurationOrDefault("REDIS_DIAL_TIMEOUT", DefaultConfig.DialTimeout),
		ReadTimeout:  utils.GetEnvDurationOrDefault("REDIS_READ_TIMEOUT", DefaultConfig.ReadTimeout),
		WriteTimeout: utils.GetEnvDurationOrDefault("REDIS_WRITE_TIMEOUT", DefaultConfig.WriteTimeout),
		PoolSize:     utils.GetEnvIntOrDefault("REDIS_POOL_SIZE", DefaultConfig.PoolSize),
		Protocol:     utils.GetEnvIntOrDefault("REDIS_PROTOCOL", DefaultConfig.Protocol),
	}

	return config, nil
}

func (*Config) BasicCheck() error {
	return nil
}
