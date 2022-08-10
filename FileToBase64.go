package utils

import (
	"bufio"
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"
)

// FileToBase64 converts a file to Base64 encoded string
func FileToBase64(filePath string) string {
	// Open file on disk.
	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	// Read entire JPG into byte slice.
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)
	return encoded
}
