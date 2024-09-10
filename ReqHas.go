package utils

import "net/http"

// ReqHas returns true if GET or POST key exists
func ReqHas(r *http.Request, key string) bool {
	if ReqGetHas(r, key) {
		return true
	}

	return ReqPostHas(r, key)
}

// ReqPostHas returns true if POST key exists
func ReqPostHas(r *http.Request, key string) bool {
	err := r.ParseForm()

	if err != nil {
		return false
	}

	_, exists := r.Form[key]
	return exists
}

// ReqGetHas returns true if GET key exists
func ReqGetHas(r *http.Request, key string) bool {
	_, exists := r.URL.Query()[key]
	return exists
}
