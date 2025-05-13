package ezex_notification

import "github.com/ezex-io/gopkg/env"

type Config struct {
	Address string
}

func LoadFromEnv() *Config {
	return &Config{
		Address: env.GetEnv[string]("EZEX_GATEWAY_NOTIFICATION_ADDRESS"),
	}
}

func (*Config) BasicCheck() error {
	return nil
}
