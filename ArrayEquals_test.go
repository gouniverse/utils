package utils

import (
	"testing"
)

func TestArrayEqualsStr(t *testing.T) {
	arr1 := []string{"one", "two"}
	arr2 := []string{"one", "two", "three"}

	if ArrayEqualsStr(arr1, arr1) == false {
		t.Error("ArrayEqualStr does not match same arrays")
	}

	if ArrayEqualsStr(arr1, arr2) {
		t.Error("ArrayEqualStr matches different arrays")
	}
}

func TestArrayEqualsInt(t *testing.T) {
	arr1 := []int{1, 2}
	arr2 := []int{1, 2, 3}

	if ArrayEqualsInt(arr1, arr1) == false {
		t.Error("ArrayEqualStr does not match same arrays")
	}

	if ArrayEqualsInt(arr1, arr2) {
		t.Error("ArrayEqualStr matches different arrays")
	}
}
