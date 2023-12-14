package utils

import (
	"net/http"
	"time"
)

func CookieRemove(w http.ResponseWriter, r *http.Request, name string) {
	secureCookie := true
	if r.TLS == nil {
		secureCookie = false // the scheme is HTTP
	}
	expiration := time.Now().Add(-365 * 24 * time.Hour)
	cookie := http.Cookie{
		Name:     name,
		Value:    "none",
		Expires:  expiration,
		HttpOnly: false,
		Secure:   secureCookie,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)
}
