package utils

import (
	"fmt"
	"strconv"
	"unsafe"
)

// ToString converts an interface to string
func ToString(v interface{}) string {
	switch v := v.(type) {
	case string:
		return v

	case []byte:
		return btos(v)

	case int:
		return strconv.Itoa(v)

	case float64:
		return strconv.FormatFloat(v, 'f', 4, 64)

	default:
		return fmt.Sprint(v)
	}
}

// btox converts bytes to string
func btos(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
