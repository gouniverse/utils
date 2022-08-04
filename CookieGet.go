package utils

import (
	"log"
	"net/http"
)

func CookieGet(r *http.Request, name string) string {
	cookie, err := r.Cookie(name)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return cookie.Value
}
