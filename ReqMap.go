package utils

import (
	"net/http"
	"strings"
)

// ReqMap returns a map from the request
func ReqMap(r *http.Request, key string) map[string]string {
	all := ReqAll(r)

	reqMap := map[string]string{}

	if all == nil {
		return reqMap
	}

	for k, v := range all {
		if strings.HasPrefix(k, key+"[") && strings.HasSuffix(k, "]") {
			if len(v) < 1 {
				reqMap[strings.TrimSuffix(strings.TrimPrefix(k, key+"["), "]")] = ""
				continue
			}

			reqMap[strings.TrimSuffix(strings.TrimPrefix(k, key+"["), "]")] = v[0]
		}
	}

	return reqMap
}
