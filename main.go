package caser

import (
	"strings"
	"unicode"
)

func CamelToSnake(str string) string  {
	var (
		words []string
		temp string
	)

	for index, char := range str {
		if index > 0 &&
			((unicode.IsUpper(char) && (unicode.IsLower(rune(str[index - 1])) || unicode.IsNumber(rune(str[index - 1])))) ||
			(unicode.IsNumber(char) && !unicode.IsNumber(rune(str[index - 1])))) {
			words = append(words, temp)

			temp = ""
		}

		temp += string(unicode.ToLower(char))
	}

	words = append(words, temp)

	return strings.Join(words, "_")
}