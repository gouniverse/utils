package utils

// ArrayMerge merges two arrays
func ArrayMerge(array ...[]interface{}) []interface{} {
	n := 0
	for _, v := range array {
		n += len(v)
	}
	outArray := make([]interface{}, 0, n)
	for _, v := range array {
		outArray = append(outArray, v...)
	}
	return outArray
}
