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

var DefaultConfig = &Config{
	ConfirmationCodeTTL:      5 * time.Minute,
	ConfirmationTemplateName: "confirmation_letter",
	ConfirmationCodeSubject:  "ezeX Confirmation Code: %s",
}

func LoadFromEnv() (*Config, error) {
	config := &Config{
		ConfirmationCodeTTL:      utils.GetEnvDurationOrDefault("AUTH_CONFIRMATION_CODE_TTL", DefaultConfig.ConfirmationCodeTTL),
		ConfirmationTemplateName: utils.GetEnvOrDefault("AUTH_CONFIRMATION_TEMPLATE", DefaultConfig.ConfirmationTemplateName),
		ConfirmationCodeSubject:  utils.GetEnvOrDefault("AUTH_CONFIRMATION_SUBJECT", DefaultConfig.ConfirmationCodeSubject),
	}

	return config, nil
}

func (*Config) BasicCheck() error {
	return nil
}
