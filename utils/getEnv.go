package utils

import (
	"os"
)

func GetEnv(key string, fallback string) string {
	env := os.Getenv(key)

	if env == "" {
		return fallback
	}

	return env
}
