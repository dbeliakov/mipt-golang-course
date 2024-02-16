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
