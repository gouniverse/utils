package utils

// Deprecated: Moved to https://github.com/gouniverse/strutils
func StrPadLeft(input string, padLength int, padString string) string {
	output := ""
	inputLen := len(input)
	if inputLen >= padLength {
		return input
	}
	ll := padLength - inputLen
	for i := 1; i <= ll; i = i + len(padString) {
		output += padString
	}
	return output + input
}
