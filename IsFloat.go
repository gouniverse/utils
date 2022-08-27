package utils

import "strconv"

// IsFloat checks if the string is a float.
func IsFloat(str string) bool {
	if IsEmpty(str) {
		return true
	}

	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}
