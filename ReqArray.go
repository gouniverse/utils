package utils

import (
	"net/http"
	"sort"
	"strings"

	"github.com/spf13/cast"
)

func ReqArray(r *http.Request, key string, defaultValue []string) []string {
	all := ReqAll(r)

	if all == nil {
		return defaultValue
	}

	isReqNumbered := false
	reqValues := []string{}
	reqKeys := []string{}

	for k, v := range all {
		isNotNumbered := k == key // key matches key, key sending full array

		if isNotNumbered {
			reqValues = append(reqValues, v...)
			break
		}

		isFullKey := k == key+`[]` // key matches key[]

		if isFullKey {
			reqValues = append(reqValues, v...)
			break
		}

		isNumbered := strings.HasPrefix(k, key+"[") && strings.HasSuffix(k, "]") // key matches key[number]

		if isNumbered {
			isReqNumbered = true
			reqKeys = append(reqKeys, k)
			if len(v) < 1 {
				reqValues = append(reqValues, "")
			} else {
				reqValues = append(reqValues, v[0])
			}
		}

	}

	if isReqNumbered {
		originalKeyIndex := map[string]int{}
		keyNumber := map[string]int{}

		// sort reqKeys and keep order of reqValues
		for i, k := range reqKeys {
			originalKeyIndex[k] = i
			keyIndex := strings.TrimSuffix(strings.TrimPrefix(k, key+"["), "]")
			keyNumber[k] = cast.ToInt(keyIndex)
		}

		sortedNumbers := []int{}

		for _, v := range keyNumber {
			sortedNumbers = append(sortedNumbers, v)
		}

		sort.Ints(sortedNumbers)

		sortedReqValues := []string{}
		for _, sortedNumber := range sortedNumbers {
			foundKey := ""
			for key, number := range keyNumber {
				if number == sortedNumber {
					foundKey = key
				}
			}
			if foundKey == "" {
				continue
			}
			sortedReqValues = append(sortedReqValues, reqValues[originalKeyIndex[foundKey]])
		}
		reqValues = sortedReqValues
	}

	return reqValues
}
