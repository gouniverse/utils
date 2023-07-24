package utils

import "net/http"

func ReqArray(r *http.Request, key string, defaultValue []string) []string {
	r.ParseForm()
	postValues := r.Form[key]

	if len(postValues) > 0 {
		return postValues
	}

	values := r.URL.Query()

	if values.Has(key) {
		return values[key]
	}

	if len(values) > 0 {
		return values[key]
	}

	return defaultValue
}
