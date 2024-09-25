package utils

import "strings"

func NL2BR(s string) string {
	return strings.ReplaceAll(s, "\n", "<br />")
}
