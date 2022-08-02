package utils

import (
	"testing"
)

func TestToString(t *testing.T) {
	if ToString(1) != "1" {
        t.Error("Conversion int to string failed")
	}
	if ToString("Hello.txt") != "Hello.txt" {
        t.Error("Conversion string to string failed")
	}
}
