package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// StrToSHA256Hash converts a string to SHA256 hash
func StrToSHA256Hash(text string) string {
	hash := sha256.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}
