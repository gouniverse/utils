package utils

// JSONEncode shortcut to ToJSON
func JSONEncode(value interface{}) (string, error) {
	return ToJSON(value)
}
