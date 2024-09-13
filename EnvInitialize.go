package utils

import (
	"log"

	"github.com/joho/godotenv"
)

// EnvInitialize initializes the environment variables
func EnvInitialize(envFilePath string) {
	if envFilePath == "" {
		envFilePath = ".env"
	}
	if FileExists(envFilePath) {
		err := godotenv.Load(envFilePath)
		if err != nil {
			log.Fatal("Error loading " + envFilePath + " file")
		}
	}
}
