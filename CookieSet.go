package utils

import (
	"net/http"
	"time"
)

func CookieSet(w http.ResponseWriter, name string, value string, seconds int) {
	expiration := time.Now().Add(time.Duration(seconds) * time.Second)
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  expiration,
		HttpOnly: false,
		Secure:   true,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)
}
