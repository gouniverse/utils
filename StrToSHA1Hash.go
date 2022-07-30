package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

// StrToSHA1Hash converts a string to SHA1 hash
func StrToSHA1Hash(text string) string {
	hash := sha1.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
