package utils

// AppAddress return the full URL and PORT for the app
func AppAddress() string {
	appAddress := Env("APP_URL") + ":" + Env("APP_PORT")
	return appAddress
}

// AppEnv returns the environment the app is running in
func AppEnv() string {
	appEnv := Env("APP_ENV")
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
	return Env("APP_NAME")
}

// AppPort return the port for the app
func AppPort() string {
	return Env("APP_PORT")
}

// AppURL return the URL for the app
func AppURL() string {
	return Env("APP_URL")
}

// DbDriver returns the driver of the database
func DbDriver() string {
	return Env("DB_DRIVER")
}

// DbHost return the host of the database
func DbHost() string {
	return Env("DB_HOST")
}

// DbPort return the port of the database
func DbPort() string {
	return Env("DB_PORT")
}

// DbDatabase return the name of the database
func DbDatabase() string {
	return Env("DB_DATABASE")
}

// DbUsername return the username of the database
func DbUsername() string {
	return Env("DB_USERNAME")
}

// DbPassword return the password for the database
func DbPassword() string {
	return Env("DB_PASSWORD")
}

// EmailFromAddress return the URL for the app
func EmailFromAddress() string {
	return Env("EMAIL_FROM_ADDRESS")
}

// EmailFromName return the URL for the app
func EmailFromName() string {
	return Env("EMAIL_FROM_NAME")
}
