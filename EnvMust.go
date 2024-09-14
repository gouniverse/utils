package utils

import (
	"log"
	"os"
)

// EnvMust retrieves the value of an environment variable, panicking if not set.
func EnvMust(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Panicf("Environment variable %s is required but not set", key)
	}
	return value
}
