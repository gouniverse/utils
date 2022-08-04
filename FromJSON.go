package utils

import "encoding/json"

func FromJSON(jsonString string, valueDefault interface{}) (interface{}, error) {
	var e interface{}

	jsonError := json.Unmarshal([]byte(jsonString), &e)

	if jsonError != nil {
		return valueDefault, jsonError
	}

	return e, nil
}
