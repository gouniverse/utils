package utils

import (
	"net/http"
	"net/url"
)

func AllRequestVariables(r *http.Request) url.Values {
	r.ParseForm()
	return r.Form
}
