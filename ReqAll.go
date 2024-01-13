package utils

import (
	"net/http"
	"net/url"
)

// ReqAll returns all request variables
func ReqAll(r *http.Request) url.Values {
	r.ParseForm()
	return r.Form
}
