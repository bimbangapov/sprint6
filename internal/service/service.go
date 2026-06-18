package service

import (
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func hasLetters(str string) bool {
	count := 0
	for _, letter := range str {
		if unicode.IsLetter(letter) {
			return true
		}

		//Проверяем только первые 100 символов
		count++
		if count == 100 {
			return false
		}

	}
	return false
}

func containsMorse(s string) bool {
	if hasLetters(s) {
		return false
	}

	return strings.ContainsAny(s, "-.")
}

func Convert(s string) string {
	if s == "" {
		return ""
	}

	if containsMorse(s) {
		return morse.ToText(s)
	}

	return morse.ToMorse(s)
}
