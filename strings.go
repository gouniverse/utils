package utils


import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"errors"
	"strconv"
	"unicode"
	"unsafe"

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

// RandStr generates random string of specified length
func RandStr(len int) string {
	buff := make([]byte, len)
	rand.Read(buff)
	str := base64.StdEncoding.EncodeToString(buff)
	
	// Base 64 can be longer than len
	return str[:len]
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

// StrToInt converts a string to Int32
func StrToInt(s string) (int, error) {
	toInt, err :=strconv.Atoi(s)
	if err != nil {
		return 0, errors.New(err.Error())
	}
	return toInt, nil
 
}


// StrToInt64 converts a string to Int64
func StrToInt64(s string) (int64, error) {
	toInit64,err  :=strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, errors.New(err.Error())
	
	}
	return toInit64, nil
}

// ToString converts an interface to string
func ToString(v interface{}) string {
	switch v := v.(type) {
	case string:
		return v

	case []byte:
		return btos(v)

	case int:
		return strconv.Itoa(v)

	case float64:
		return strconv.FormatFloat(v, 'f', 4, 64)

	default:
		return fmt.Sprint(v)
	}
}

// btox converts bytes to string
func btos(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// ln returns new line as bytes
func ln() []byte {
	return stob("\n")
}

// stob converts string to bytes
func stob(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// pxToString converts int to string (i.e. 1px)
func pxToString(px int) string {
	return strconv.Itoa(px) + "px"
}
