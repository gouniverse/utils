package utils

import "strings"

func StrContainsOnlySpecifiedCharacters(str string, chars string) bool {
	nonCharacter := func(c rune) bool { return !strings.ContainsAny(chars, string(c)) }
	words := strings.FieldsFunc(str, nonCharacter)
	return str == strings.Join(words, "")
}
