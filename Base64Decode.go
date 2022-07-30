package utils

import "encoding/base64"

// Base64Encode decodes a string from Base64
func Base64Decode(src string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(src)
}
