package utils

import (
	"net/http"
	"strings"
)

func ReqArray(r *http.Request, key string, defaultValue []string) []string {
	all := ReqAll(r)

	reqArray := []string{}

	if all == nil {
		return defaultValue
	}

	for k, v := range all {
		if !strings.HasPrefix(k, key+"[") || !strings.HasSuffix(k, "]") {
			continue
		}

		if len(v) < 1 {
			reqArray = append(reqArray, "")
			continue
		}

		reqArray = append(reqArray, v[0])
	}

	return reqArray
}
