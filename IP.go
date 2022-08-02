package utils

import (
	"net"
	"net/http"
	"strings"
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
