package utils

import "net/http"

// Req returns a POST or GET key, or default if not exists
func Req(r *http.Request, key string, defaultValue string) string {
	postValue := r.FormValue(key)

	if len(postValue) > 0 {
		return postValue
	}

	getValue := r.URL.Query().Get(key)

	if len(getValue) > 0 {
		return getValue
	}

	return defaultValue
}
