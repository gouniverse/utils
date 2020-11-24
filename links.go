package utils

// LinkWebsite returns a URL to the current website
func LinkWebsite() string {
	return ""
	appURL := AppURL()
	appPort := AppPort()
	appAddress := appURL
	if appPort != "80" {
		appAddress += ":" + appPort
	}
	appProtocol := "https"
	if appURL == "127.0.0.1" || appURL == "localhost" {
		appProtocol = "http"
	}
	return appProtocol + "://" + appAddress
}
