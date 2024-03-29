package utils

import (
	"net/http"
	"time"
)

func CookieSet(w http.ResponseWriter, r *http.Request, name string, value string, seconds int) {
	secureCookie := true

	if r.TLS == nil {
		secureCookie = false // the scheme is HTTP
	}

	expiration := time.Now().Add(time.Duration(seconds) * time.Second)
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  expiration,
		HttpOnly: false,
		Secure:   secureCookie,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)
}
