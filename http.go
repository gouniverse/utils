package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gouniverse/api"
)

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

// Respond returns an API response as JSON
func RespondJSON(w http.ResponseWriter, response api.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(api.Success("API Working on " + environment))
	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}
