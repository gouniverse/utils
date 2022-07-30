package utils

import (
	"strconv"
)

// IsNumeric checks if a string is numeric
func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
