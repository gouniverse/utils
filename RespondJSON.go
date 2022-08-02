package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gouniverse/api"
)

// RespondJSON returns an API response as JSON
func RespondJSON(w http.ResponseWriter, response api.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}
