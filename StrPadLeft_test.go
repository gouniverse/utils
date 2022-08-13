package utils

import "testing"

func TestStrPadLeft(t *testing.T) {
	str := "1"
	expected := "01"

	result := StrPadLeft(str, 2, "0")

	if result != expected {
		t.Errorf("StrPadLeft does not return expected result '01', but '%s'", result)
	}
}
