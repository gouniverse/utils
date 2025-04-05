package utils

// Deprecated: ArrayEquals checks whether 2 string arrays are the same
// use Equals from github.com/dracory/arr
func ArrayEqualsStr(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Deprecated: ArrayEquals checks whether 2 string arrays are the same
// use Equals from github.com/dracory/arr
func ArrayEqualsInt(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
