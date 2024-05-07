package lib

import (
	"strings"
	"unicode"
)

func Encode() {

}

func prepareText(str string) string {
	var builder strings.Builder

	for _, ch := range str {
		if unicode.IsUpper(ch) {
			builder.WriteRune('!')
			builder.WriteRune(unicode.ToLower(ch))
		} else {
			builder.WriteRune(ch)
		}
	}

	return builder.String()
}
