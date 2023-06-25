package utils

import (
	"bytes"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// StrSlugify replaces each run of characters which are not ASCII letters or numbers
// with the Replacement character, except for leading or trailing runs. Letters
// will be stripped of diacritical marks and lowercased. Letter or number
// codepoints that do not have combining marks or a lower-cased variant will be
// passed through unaltered.
func StrSlugify(s string, replaceWith rune) string {
	s = cyrillicToLatin(s)

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

func cyrillicToLatin(s string) string {
	var sub = map[rune]string{
		'А': "A",
		'Б': "B",
		'В': "V",
		'Г': "G",
		'Д': "D",
		'Е': "E",
		'Ж': "Zh",
		'З': "Z",
		'И': "I",
		'Й': "Y",
		'К': "K",
		'Л': "L",
		'М': "M",
		'Н': "N",
		'О': "O",
		'П': "P",
		'Р': "R",
		'С': "S",
		'Т': "T",
		'У': "U",
		'Ф': "F",
		'Х': "H",
		'Ц': "Ts",
		'Ч': "Ch",
		'Ш': "Sh",
		'Щ': "Sh",
		'Ъ': "A",
		'Ь': "Y",
		'Ю': "Yu",
		'Я': "Ya",
		'а': "a",
		'б': "b",
		'в': "v",
		'г': "g",
		'д': "d",
		'е': "e",
		'ж': "zh",
		'з': "z",
		'и': "i",
		'й': "y",
		'к': "k",
		'л': "l",
		'м': "m",
		'н': "n",
		'о': "o",
		'п': "p",
		'р': "r",
		'с': "s",
		'т': "t",
		'у': "u",
		'ф': "f",
		'х': "h",
		'ц': "ts",
		'ч': "ch",
		'ш': "sh",
		'щ': "sht",
		'ъ': "a",
		'ь': "y",
		'ю': "yu",
		'я': "ya",
	}

	var buf bytes.Buffer
	for _, c := range s {
		if d, ok := sub[c]; ok {
			buf.WriteString(d)
		} else {
			buf.WriteRune(c)
		}
	}
	return buf.String()
}
