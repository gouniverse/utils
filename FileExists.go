package utils

import "os"

// FileExists checks if a file exists
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)

	return !os.IsNotExist(err)
}
