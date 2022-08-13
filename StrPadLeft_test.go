package utils

import "testing"

func TestStrPadLeft(t *testing.T) {
	str := "1"
	expected := "01"

	result := StrPadLeft(str, 2, "0")

	if result != expected {
		t.Error("StrPadLeft does not return expected result, but %s", result)
	}
}
