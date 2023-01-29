package utils

import "os"

// FileExists checks if a file exists
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		return false
	}

	return true
}
