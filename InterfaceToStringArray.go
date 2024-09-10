package utils

import (
	"fmt"
	"reflect"
)

// InterfaceToStringArray converts an interface to String array
func InterfaceToStringArray(v interface{}) []string {
	if v == nil {
		return []string{}
	}
	val := reflect.ValueOf(v)
	fmt.Println(val.Kind())

	if val.Kind() != reflect.Slice {
		return []string{}
	}

	stringArrayInterface := v.([]string)
	stringArray := make([]string, len(stringArrayInterface))
	copy(stringArray, stringArrayInterface)

	return stringArray
}
