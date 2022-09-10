package utils

import (
	"testing"
)

func TestStrContainsOnlySpecifiedChars(t *testing.T) {
	chars := "ABCDEF"

	if !StrContainsOnlySpecifiedCharacters("ABC", chars) {
		t.Error("StrContainsOnlySpecifiedChars MUST return true for ABC contains only chars: " + chars)
	}

	if StrContainsOnlySpecifiedCharacters("XYZ", chars) {
		t.Error("StrContainsOnlySpecifiedChars MUST return false for XYZ contains only chars: " + chars)
	}

	if StrContainsOnlySpecifiedCharacters("XAZ", chars) {
		t.Error("StrContainsOnlySpecifiedChars MUST return false for XAZ contains only chars: " + chars)
	}
}
