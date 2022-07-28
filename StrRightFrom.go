package utils

import "strings"

// StrRightFrom returns the substring on the left side of the needle
func StrRightFrom(str, needle string) string {
	_, right, isFound := strings.Cut(str, needle)

	if !isFound {
		return ""
	}

	return right
}
