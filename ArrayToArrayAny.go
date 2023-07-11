package utils

// ArrayToArrayAny converts an array of any type to array of interfaces
func ArrayToArrayAny[T any](arr []T) []any {
	var ret []any
	for _, v := range arr {
		ret = append(ret, v)
	}
	return ret
}
