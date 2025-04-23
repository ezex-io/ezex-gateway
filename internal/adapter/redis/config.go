package redis

import (
	"time"

	"github.com/ezex-io/gopkg/env"
)

type Config struct {
	Address      string
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
		Address:      env.GetEnv[string]("EZEX_GATEWAY_REDIS_HOST", env.WithDefault("localhost:6379")),
		DB:           env.GetEnv[int]("EZEX_GATEWAY_REDIS_DB", env.WithDefault("0")),
		Password:     env.GetEnv[string]("EZEX_GATEWAY_REDIS_PASSWORD"),
		DialTimeout:  env.GetEnv[time.Duration]("EZEX_GATEWAY_REDIS_DIAL_TIMEOUT", env.WithDefault("5s")),
		ReadTimeout:  env.GetEnv[time.Duration]("EZEX_GATEWAY_REDIS_READ_TIMEOUT", env.WithDefault("5s")),
		WriteTimeout: env.GetEnv[time.Duration]("EZEX_GATEWAY_REDIS_WRITE_TIMEOUT", env.WithDefault("5s")),
		PoolSize:     env.GetEnv[int]("EZEX_GATEWAY_REDIS_POOL_SIZE", env.WithDefault("10")),
		Protocol:     env.GetEnv[int]("EZEX_GATEWAY_REDIS_PROTOCOL", env.WithDefault("3")),
	}

	return config, nil
}

func (*Config) BasicCheck() error {
	return nil
}
