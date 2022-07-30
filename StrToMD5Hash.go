package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// StrToMD5Hash converts a string to MD5 hash
func StrToMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
