package palindrome

import (
	"unicode"
)

func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if !unicode.IsLetter(r) {
			panic("123")
		}
		letters = append(letters, unicode.ToLower(r))
	}

	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
