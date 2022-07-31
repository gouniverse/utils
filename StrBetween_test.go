package utils

import "testing"

func TestStrBetween(t *testing.T) {
	needle1 := "needle1"
	needle2 := "needle2"
	str := "every " + needle1 + " cloud has a silver " + needle2 + " lining"
	expected := " cloud has a silver "

	result, isFound := StrBetween(str, needle1, needle2)

	if result != expected {
		t.Error("StrBetween does not return expected result")
	}

	if !isFound {
		t.Error("StrBetween does not find the needles")
	}
}
