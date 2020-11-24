package utils

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

func Env(key string) string {
    val := os.Getenv(key)
    return val
}
