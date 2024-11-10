package utils

import (
	"net/http"
	"strings"
)

func ReqArray(r *http.Request, key string, defaultValue []string) []string {
	all := ReqAll(r)

	if all == nil {
		return defaultValue
	}

	reqArray := []string{}

	for k, v := range all {
		isFullKey := k == key+`[]`

		if isFullKey {
			reqArray = append(reqArray, v...)
			break
		}

		isSameKey := strings.HasPrefix(k, key+"[") && strings.HasSuffix(k, "]")

		if isSameKey {
			if len(v) < 1 {
				reqArray = append(reqArray, "")
			} else {
				reqArray = append(reqArray, v[0])
			}
		}

		isNotNumbered := k == key

		if isNotNumbered {
			reqArray = append(reqArray, v...)
			break
		}
	}

	return reqArray
}
