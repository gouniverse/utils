package utils

// JSONDecode shortcode for FromJSON
func JSONDecode(jsonString string, valueDefault interface{}) (interface{}, error) {
	return FromJSON(jsonString, valueDefault)
}
