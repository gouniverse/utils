package utils

import "strings"

// StrBetween returns the string between two needles
//
// Deprecated: Moved to https://github.com/gouniverse/strutils
func StrBetween(str string, startNeedle string, endNeedle string) (result string, found bool) {
	s := strings.Index(str, startNeedle)
	if s == -1 {
		return result, false
	}

	newStr := str[s+len(startNeedle):]
	e := strings.Index(newStr, endNeedle)
	if e == -1 {
		return result, false
	}
	result = newStr[:e]
	return result, true
}
