package utils

import (
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {
    if FileExists("Hello.txt") {
        t.Error("Non-existing file exists")
	}

	FilePutContents("Hello.txt","Hello",os.FileMode(os.O_APPEND))
	
	if FileExists("Hello.txt") ==false{
        t.Error("File DOES NOT exist")
	}
}
