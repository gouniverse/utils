package utils

// Deprecated. ArrayReverse creates a new reversed array
// use Reverse from github.com/dracory/base/arr
func ArrayReverse[T any](arr []T) []T {
	newArr := make([]T, len(arr))

	next := 0
	for i := len(arr) - 1; i > -1; i-- {
		newArr[next] = arr[i]
		next++
	}

	return newArr
}
