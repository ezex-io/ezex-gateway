package auth

import "time"

type Config struct {
	ConfirmationCodeTTL      time.Duration `yaml:"confirmation_code_ttl"`
	ConfirmationTemplateName string        `yaml:"confirmation_template_name"`
	ConfirmationCodeSubject  string        `yaml:"confirmation_code_subject"`
}

var DefaultConfig = &Config{
	ConfirmationCodeTTL:      5 * time.Minute,
	ConfirmationTemplateName: "confirmation_letter",
	ConfirmationCodeSubject:  "ezeX Confirmation Code: %s",
}

func (*Config) BasicCheck() error {
	return nil
}
