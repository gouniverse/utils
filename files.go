package utils

import (
	"bufio"
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// FileExists checks if a file exists
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// FileGetContents reads entire file into a string
func FileGetContents(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	return string(data), err
}


// FilePutContents adds content to file
func FilePutContents(filename string, data string, mode os.FileMode) error {
	return ioutil.WriteFile(filename, []byte(data), mode)
}

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

// ImgToBase64Url converts an image file to Base64 encoded URL string
func ImgToBase64Url(filePath string) string {
	// Read the entire file into a byte slice
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	base64Encoding := ""

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/bmp":
		base64Encoding += "data:image/bmp;base64,"
	case "image/gif":
		base64Encoding += "data:image/gif;base64,"
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	case "image/webp":
		base64Encoding += "data:image/webp;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += FileToBase64(filePath)
	return base64Encoding
}

// Unlink deletes a file
func Unlink(filename string) error {
	return os.Remove(filename)
}
