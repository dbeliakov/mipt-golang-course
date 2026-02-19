package caesar_test

import (
	"github.com/dbeliakov/mipt-golang-course/tasks/02/caesar"
	"testing"
)

func TestCipher(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		shift  int
		expect string
	}{
		{"BasicLowerCase", "abc", 3, "def"},
		{"BasicUpperCase", "ABC", 3, "DEF"},
		{"WrapAround", "xyz", 3, "abc"},
		{"FullSentence", "Hello, World!", 3, "Khoor, Zruog!"},
		{"NegativeShift", "def", -3, "abc"},
		{"ZeroShift", "abc", 0, "abc"},
		{"NonAlphabeticCharacters", "123", 3, "123"},
		{"CyrillicLower", "абв", 3, "где"},
		{"CyrillicUpper", "АБВ", 3, "ГДЕ"},
		{"CyrillicWrapLower", "эюя", 3, "абв"},
		{"CyrillicWrapUpper", "ЭЮЯ", 3, "АБВ"},
		{"CyrillicNegative", "где", -3, "абв"},
		{"CyrillicMixed", "abc абв", 1, "bcd бвг"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := caesar.Cipher(c.input, c.shift)
			if got != c.expect {
				t.Errorf("caesarCipher(%q, %d) == %q, want %q", c.input, c.shift, got, c.expect)
			}
		})
	}

}
