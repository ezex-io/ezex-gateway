package auth

import (
	"time"

	"github.com/ezex-io/ezex-gateway/internal/utils"
)

type Config struct {
	ConfirmationCodeTTL      time.Duration
	ConfirmationTemplateName string
	ConfirmationCodeSubject  string
}

func LoadFromEnv() (*Config, error) {
	config := &Config{
		ConfirmationCodeTTL: utils.GetEnvDurationOrDefault(
			"EZEX_AUTH_CONFIRMATION_CODE_TTL",
			5*time.Minute,
		),
		ConfirmationTemplateName: utils.GetEnvOrDefault(
			"EZEX_AUTH_CONFIRMATION_TEMPLATE",
			"confirmation_letter",
		),
		ConfirmationCodeSubject: utils.GetEnvOrDefault(
			"EZEX_AUTH_CONFIRMATION_SUBJECT",
			"ezeX Confirmation Code: %s",
		),
	}

	return config, nil
}

func (*Config) BasicCheck() error {
	return nil
}
