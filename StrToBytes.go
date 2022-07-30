package utils

import "unsafe"

// StrToBytes converts string to bytes
func StrToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
