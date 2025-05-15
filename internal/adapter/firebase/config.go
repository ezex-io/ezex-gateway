package firebase

import "github.com/ezex-io/gopkg/env"

type Config struct {
	ProjectID string
	APIKey    string
}

func LoadFromEnv() *Config {
	return &Config{
		APIKey:    env.GetEnv[string]("EZEX_GATEWAY_FIREBASE_API_KEY"),
		ProjectID: env.GetEnv[string]("EZEX_GATEWAY_FIREBASE_PROJECT_ID"),
	}
}

func (*Config) BasicCheck() error {
	// Add validation if needed
	return nil
}
