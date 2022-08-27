package utils

import (
	"testing"
)

func TestToInt(t *testing.T) {
	theint, err := ToInt("1")

	if err != nil {
		t.Error("Conversion \"1\" to int failed")
	}

	if theint != 1.0 {
		t.Errorf("Result is not 1 but %d", theint)
	}
}
