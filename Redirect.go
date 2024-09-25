package utils

import "net/http"

func Redirect(w http.ResponseWriter, r *http.Request, url string) string {
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return ""
}
