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
	if appEnv == "testing" || appEnv == "test" {
		return "testing"
	}
	return appEnv
}

// AppInDevelopment return whether app is in development
func AppInDevelopment() bool {
	return AppEnv() == "development"
}

// AppInProduction return whether app is in production / live
func AppInLive() bool {
	return AppEnv() == "live"
}

// AppInProduction return whether app is in production / live
func AppInProduction() bool {
	return AppEnv() == "live"
}

// AppInTesting return whether app is being tested
func AppInTesting() bool {
	return AppEnv() == "testing"
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

// DbDriver returns the driver of the database
func DbDriver() string {
	appURL := os.Getenv("DB_DRIVER")
	return appURL
}

// DbHost return the host of the database
func DbHost() string {
	appURL := os.Getenv("DB_HOST")
	return appURL
}

// DbPort return the port of the database
func DbPort() string {
	appURL := os.Getenv("DB_PORT")
	return appURL
}

// DbDatabase return the name of the database
func DbDatabase() string {
	appURL := os.Getenv("DB_DATABASE")
	return appURL
}

// DbUsername return the username of the database
func DbUsername() string {
	appURL := os.Getenv("DB_USERNAME")
	return appURL
}

// DbPassword return the password for the database
func DbPassword() string {
	appURL := os.Getenv("DB_PASSWORD")
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
