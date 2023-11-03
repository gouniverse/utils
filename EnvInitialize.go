package utils

import (
	"log"

	"github.com/joho/godotenv"
)

// EnvInitialize initializes the environment variables
func EnvInitialize() {
	if FileExists(".env") {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}
