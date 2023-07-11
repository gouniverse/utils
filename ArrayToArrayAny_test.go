package utils

import (
	"reflect"
	"testing"
)

func TestArrayToArrayAny(t *testing.T) {
	arr := []string{"one", "two", "three"}
	newArr := ArrayToArrayAny(arr)

	if !reflect.DeepEqual(newArr, []any{"one", "two", "three"}) {
		t.Error("The 2 arrays must be equal:", arr, newArr)
	}
}
