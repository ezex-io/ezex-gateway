package ezex_users

import "github.com/ezex-io/gopkg/env"

type Config struct {
	Address string
}

func LoadFromEnv() *Config {
	return &Config{
		Address: env.GetEnv[string]("EZEX_GATEWAY_USERS_ADDRESS"),
	}
}

func (*Config) BasicCheck() error {
	return nil
}
