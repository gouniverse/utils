package utils

import "strings"

// StrLeftFrom returns the substring on the left side of the needle
func StrLeftFrom(str, needle string) string {
	left, _, isFound := strings.Cut(str, needle)

	if !isFound {
		return ""
	}

	return left
}
