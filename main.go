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
		} else if index > 0 && (unicode.IsUpper(char) || (unicode.IsNumber(char) && !unicode.IsNumber(prevChar))) {
			words = append(words, temp)

			temp = ""
		}

		temp += string(unicode.ToLower(char))

		prevChar = char
	}

	words = append(words, temp)

	return strings.Join(words, "_")
}