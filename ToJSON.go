package utils

import "encoding/json"

func ToJSON(value interface{}) (string, error) {
	jsonValue, jsonError := json.Marshal(value)
	if jsonError != nil {
		return "", jsonError
	}

	return string(jsonValue), nil
}
