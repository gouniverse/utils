package utils

import "bytes"

// AddSlashes returns a string with backslashes added before characters that need to be escaped.
//
// These characters are:
// single quote (')
// double quote (")
// backslash (\)
//
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
