package utils

import "reflect"

// ArrayContains checks whether an array contains the specified value
func ArrayContains(array interface{}, val interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

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

// IsStringArrayEqual checks whether 2 string arrays are the same
func IsStringArrayEqual(a, b []string) bool {
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
