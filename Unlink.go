package utils

import (
	"os"
)

// Unlink deletes a file
func Unlink(filename string) error {
	return os.Remove(filename)
}
