package utils

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	mrand "math/rand"
	"strconv"
	"time"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// AddSlashes addslashes()
func AddSlashes(str string) string {
	var buf bytes.Buffer
	for _, char := range str {
		switch char {
		case '\'', '"', '\\':
			buf.WriteRune('\\')
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

// Base64Encode decodes a string from Base64
func Base64Decode(src string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(src)
}

// Base64Encode encodes a string to Base64
func Base64Encode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

// Deprecated: RandStr is deprecated, new code should use StrRandom instead.
// RandStr generates random string of specified length
func RandStr(length int) string {
	buff := make([]byte, length)
	rand.Read(buff)
	str := base64.StdEncoding.EncodeToString(buff)

	// Base 64 can be longer than len
	return str[:length]
}

// Deprecated: RandStrFromGamma is deprecated, new code should use StrRandomFromGamma instead.
// RandStrFromGamma generates random string of specified length with the characters specified in the gamma string
func RandStrFromGamma(length int, gamma string) string {
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

// Slugify replaces each run of characters which are not ASCII letters or numbers
// with the Replacement character, except for leading or trailing runs. Letters
// will be stripped of diacritical marks and lowercased. Letter or number
// codepoints that do not have combining marks or a lower-cased variant will be
// passed through unaltered.
func Slugify(s string, replaceWith rune) string {
	// The "safe" set of characters.
	alphanum := &unicode.RangeTable{
		R16: []unicode.Range16{
			{0x0030, 0x0039, 1}, // 0-9
			{0x0041, 0x005A, 1}, // A-Z
			{0x0061, 0x007A, 1}, // a-z
		},
	}
	// Characters in these ranges will be ignored.
	nop := []*unicode.RangeTable{
		unicode.Mark,
		unicode.Sk, // Symbol - modifier
		unicode.Lm, // Letter - modifier
		unicode.Cc, // Other - control
		unicode.Cf, // Other - format
	}
	buf := make([]rune, 0, len(s))
	replacement := false

	for _, r := range norm.NFKD.String(s) {
		switch {
		case unicode.In(r, alphanum):
			buf = append(buf, unicode.ToLower(r))
			replacement = true
		case unicode.IsOneOf(nop, r):
			// skip
		case replacement:
			buf = append(buf, replaceWith)
			replacement = false
		}
	}

	// Strip trailing Replacement byte
	if i := len(buf) - 1; i >= 0 && buf[i] == replaceWith {
		buf = buf[:i]
	}

	return string(buf)
}

// ln returns new line as bytes
// func ln() []byte {
// 	return StrToBytes("\n")
// }

// pxToString converts int to string (i.e. 1px)
func pxToString(px int) string {
	return strconv.Itoa(px) + "px"
}
