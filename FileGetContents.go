package utils

import "os"

// FileGetContents reads entire file into a string
func FileGetContents(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	return string(data), err
}
