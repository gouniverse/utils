package utils

import (
	"os"
)

// FilePutContents adds content to file
func FilePutContents(filename string, data string, mode os.FileMode) error {
	return os.WriteFile(filename, []byte(data), mode)
}
