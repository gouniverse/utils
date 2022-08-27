package utils

import "strconv"

// IsInt checks if the string is an integer. Empty string is valid.
func IsInt(str string) bool {
	if IsEmpty(str) {
		return true
	}

	_, err := strconv.ParseInt(str, 10, 64)
	return err == nil
}
