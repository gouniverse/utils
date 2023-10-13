package utils

import (
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {
	filePath := "FileExistsTest.txt"
	
    if FileExists(filePath) {
		Unlink(filePath)
	}
	
	if FileExists(filePath) {
        t.Error("Non-existing file exists")
	}

	FilePutContents(filePath,"Hello",os.FileMode(os.O_APPEND))
	
	if FileExists(filePath) == false{
        t.Error("File DOES NOT exist")
	}
}


