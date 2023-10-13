package utils

import (
	"os"
	"testing"
)

func TestUnlink(t *testing.T) {
	filePath := "UnlinkTest.txt"

	if FileExists(filePath) {
		Unlink(filePath)
	}

	if FileExists(filePath) {
		t.Error("Non-existing file exists")
	}

	FilePutContents(filePath, "Hello", os.FileMode(os.O_APPEND))

	if FileExists(filePath) == false {
		t.Error("File DOES NOT exist")
	}

	Unlink(filePath)

	if FileExists(filePath) {
		t.Error("Unlink failed")
	}

}
