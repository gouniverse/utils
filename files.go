package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

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
