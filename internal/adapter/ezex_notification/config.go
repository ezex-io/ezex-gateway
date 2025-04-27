package ezex_notification

import "github.com/ezex-io/gopkg/env"

type Config struct {
	Address string
}

func LoadFromEnv() (*Config, error) {
	config := &Config{
		Address: env.GetEnv[string]("EZEX_GATEWAY_NOTIFICATION_ADDRESS"),
	}

	return config, nil
}

func (*Config) BasicCheck() error {
	return nil
}
