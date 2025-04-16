package utils

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// GetEnvOrDefault returns the value of the environment variable or the default value if not set
func GetEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// GetEnvIntOrDefault returns the integer value of the environment variable or the default value if not set
func GetEnvIntOrDefault(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// GetEnvBoolOrDefault returns the boolean value of the environment variable or the default value if not set
func GetEnvBoolOrDefault(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		return value == "true"
	}
	return defaultValue
}

// GetEnvDurationOrDefault returns the duration value of the environment variable or the default value if not set
func GetEnvDurationOrDefault(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}

// GetEnvSliceOrDefault returns the slice of strings from the environment variable or the default value if not set
func GetEnvSliceOrDefault(key string, defaultValue []string) []string {
	if value := os.Getenv(key); value != "" {
		return strings.Split(value, ",")
	}
	return defaultValue
}
