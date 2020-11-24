package utils

import (
	"os"
)

// AppAddress return the full URL and PORT for the app
func AppAddress() string {
	appAddress := os.Getenv("APP_URL") + ":" + os.Getenv("APP_PORT")
	return appAddress
}

// AppEnv returns the environment the app is running in
func AppEnv() string {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "live" || appEnv == "production" {
		return "live"
	}
	if appEnv == "development" || appEnv == "dev" {
		return "development"
	}
	return appEnv
}

// AppName return the name for the app
func AppName() string {
	appName := os.Getenv("APP_NAME")
	return appName
}

// AppPort return the port for the app
func AppPort() string {
	appPort := os.Getenv("APP_PORT")
	return appPort
}

// AppURL return the URL for the app
func AppURL() string {
	appURL := os.Getenv("APP_URL")
	return appURL
}

// EmailFromAddress return the URL for the app
func EmailFromAddress() string {
	emailFromAddress := os.Getenv("EMAIL_FROM_ADDRESS")
	return emailFromAddress
}

// EmailFromName return the URL for the app
func EmailFromName() string {
	emailFromName := os.Getenv("EMAIL_FROM_NAME")
	return emailFromName
}

// Env returns the value for an environment key
func Env(key string) string {
    val := os.Getenv(key)
    return val
}
