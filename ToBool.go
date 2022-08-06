package utils

// ToBool converts a string with common names ("yes", "true", "1") to boolean
func ToBool(str string) bool {
	if str == "1" {
		return true
	}
	if str == "yes" {
		return true
	}
	if str == "true" {
		return true
	}
	return false
}
