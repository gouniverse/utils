package utils

import (
	"encoding/json"
	"net"
	"net/http"
	"strings"

	"github.com/gouniverse/api"
)

// IP gets the IP address for the user
func IP(r *http.Request) string {
	//Get IP from the X-REAL-IP header
	realIP := r.Header.Get("X-REAL-IP")
	if realIP != "" {
		return realIP
	}
	
	//Get IP from X-FORWARDED-FOR header
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		splitIps := strings.Split(forwarded, ",")
		for _, ip := range splitIps {
		    return ip
		}
		return forwarded
	}
	
	//Get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return ""
	}

	return ip
}

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

// RespondJSON returns an API response as JSON
func RespondJSON(w http.ResponseWriter, response api.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(api.Success("API Working on " + environment))
	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}
