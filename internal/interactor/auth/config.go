package auth

import (
	"time"

	"github.com/ezex-io/gopkg/env"
)

type Config struct {
	ConfirmationCodeTTL      time.Duration
	ConfirmationTemplateName string
	ConfirmationCodeSubject  string
}

func LoadFromEnv() *Config {
	return &Config{
		ConfirmationCodeTTL: env.GetEnv[time.Duration]("EZEX_GATEWAY_AUTH_CONFIRMATION_CODE_TTL",
			env.WithDefault("5m")),
		ConfirmationTemplateName: env.GetEnv[string]("EZEX_GATEWAY_AUTH_CONFIRMATION_TEMPLATE",
			env.WithDefault("confirmation_letter")),
		ConfirmationCodeSubject: env.GetEnv[string]("EZEX_GATEWAY_AUTH_CONFIRMATION_SUBJECT",
			env.WithDefault("ezeX Confirmation Code: %s")),
	}
}

func (*Config) BasicCheck() error {
	return nil
}
