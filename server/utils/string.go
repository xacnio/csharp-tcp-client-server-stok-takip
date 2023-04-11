package utils

import (
	"strings"
)

func TurkishToEnglish(input string) string {
	turkishChars := []rune("ÇçĞğİıÖöŞşÜü")
	englishChars := []rune("CcGgIiOoSsUu")
	for i, c := range turkishChars {
		input = strings.Replace(input, string(c), string(englishChars[i]), -1)
	}
	return input
}
