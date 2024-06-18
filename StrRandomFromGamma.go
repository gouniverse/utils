package utils

import (
	mrand "math/rand"
	"time"
)

// StrRandomFromGamma generates random string of specified length with the characters specified in the gamma string
//
// Deprecated: Moved to https://github.com/gouniverse/strutils
func StrRandomFromGamma(length int, gamma string) string {
	inRune := []rune(gamma)
	out := ""

	for i := 0; i < length; i++ {
		mrand.Seed(int64(i*9876543) + time.Now().UnixNano())
		randomIndex := mrand.Intn(len(inRune))
		pick := inRune[randomIndex]
		out += string(pick)
	}

	return out
}
