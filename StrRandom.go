package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// StrRandom generates random string of specified length
func StrRandom(length int) string {
	buff := make([]byte, length)
	rand.Read(buff)
	str := base64.StdEncoding.EncodeToString(buff)

	// Base 64 can be longer than len
	return str[:length]
}
