package caser

import (
	"strings"
	"unicode"
)

func CamelToSnake(str string) string  {
	var (
		words []string
		temp string
		prevChar rune
	)

	for index, char := range str {
		if index == 0 {
			char = unicode.ToLower(char)
		}

		if unicode.IsUpper(char) ||
			(unicode.IsNumber(char) && prevChar != 0 && !unicode.IsNumber(prevChar)) ||
			(unicode.IsLetter(char) && prevChar != 0 && unicode.IsNumber(prevChar)) {
			words = append(words, temp)

			temp = ""
		}

		temp += string(unicode.ToLower(char))

		prevChar = char
	}

	words = append(words, temp)

	return strings.Join(words, "_")
}