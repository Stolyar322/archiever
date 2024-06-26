package lib

import (
	"strings"
	"unicode"
)

type encodingTable map[rune]string

func Encode() {

}

func encodeBin(encodeStr string) string {
	builder := strings.Builder{}

	for _, symbol := range encodeStr {
		builder.WriteString(bin(symbol))
	}

	return builder.String()
}

func bin(symbol rune) string {
	table := getEncodingTable()
	code, ok := table[symbol]
	if !ok {
		panic("Symbol " + string(symbol) + " is not exist at encoding table")
	}

	return code
}

func getEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		't': "1001",
		'n': "10000",
		's': "0101",
		'r': "01000",
		'd': "00101",
		'!': "001000",
		'c': "000101",
		'm': "000011",
		'g': "0000100",
		'b': "0000010",
		'v': "00000001",
		'k': "0000000001",
		'q': "000000000001",
		'e': "101",
		'o': "10001",
		'a': "011",
		'i': "01001",
		'h': "0011",
		'l': "001001",
		'u': "00011",
		'f': "000100",
		'p': "0000101",
		'w': "0000011",
		'y': "0000001",
		'j': "000000001",
		'x': "00000000001",
		'z': "000000000000",
	}
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
