package utils

import (
	"testing"
)

func TestArrayReverseStr(t *testing.T) {
	arr := []string{"one", "two", "three"}
	newArr := ArrayReverse(arr)

	if len(newArr) != len(arr) {
		t.Errorf("Length of 2 arrays must match, but found %d and %d", len(arr), len(newArr))
	}

	if newArr[0] != "three" {
		t.Errorf("New array at index 0, value is not 'three' but found %s", newArr[0])
	}

	if newArr[1] != "two" {
		t.Errorf("New array at index 1, value is not 'two' but found %s", newArr[1])
	}

	if newArr[2] != "one" {
		t.Errorf("New array at index 2, value is not 'one' but found %s", newArr[2])
	}

	if arr[0] != "one" {
		t.Errorf("Old array at index 0, value is not 'one' but found %s", arr[0])
	}

	if arr[1] != "two" {
		t.Errorf("Old array at index 1, value is not 'two' but found %s", arr[1])
	}

	if arr[2] != "three" {
		t.Errorf("Old array at index 2, value is not 'two' but found %s", arr[1])
	}
}

func TestArrayReverseInt(t *testing.T) {
	arr := []int{1, 2, 3}
	newArr := ArrayReverse(arr)

	if len(newArr) != len(arr) {
		t.Errorf("Length of 2 arrays must match, but found %d and %d", len(arr), len(newArr))
	}

	if newArr[0] != 3 {
		t.Errorf("New array at index 0, value is not 3 but found %d", newArr[0])
	}

	if newArr[1] != 2 {
		t.Errorf("New array at index 1, value is not 2 but found %d", newArr[1])
	}

	if newArr[2] != 1 {
		t.Errorf("New array at index 2, value is not 1 but found %d", newArr[2])
	}

	if arr[0] != 1 {
		t.Errorf("Old array at index 0, value is not 1 but found %d", arr[0])
	}

	if arr[1] != 2 {
		t.Errorf("Old array at index 1, value is not 2 but found %d", arr[1])
	}

	if arr[2] != 3 {
		t.Errorf("Old array at index 2, value is not 3 but found %d", arr[1])
	}
}
