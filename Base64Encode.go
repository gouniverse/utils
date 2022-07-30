package utils

import "encoding/base64"

// Base64Encode encodes a string to Base64
func Base64Encode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}
