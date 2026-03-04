package rot13_test

import (
	"testing"

	"github.com/dbeliakov/mipt-golang-course/tasks/02/rot13"
)

func TestRot13(t *testing.T) {
	cases := []struct {
		name      string
		hexInput  string
		expectHex string
		expectErr bool
	}{
		{"Basic", "48656c6c6f", "5572797962", false},
		{"LowerCase", "616263", "6e6f70", false},
		{"UpperCase", "414243", "4e4f50", false},
		{"WithSpaces", "48656c6c6f2c20576f726c6421", "55727979622c204a6265797121", false},
		{"NonAlphabetic", "313233", "313233", false},
		{"Empty", "", "", false},
		{"InvalidHex", "xyz", "", true},
		{"OddLengthHex", "48656c6c6", "", true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := rot13.Rot13(c.hexInput)
			if c.expectErr {
				if err == nil {
					t.Errorf("Rot13(%q) expected error, got nil", c.hexInput)
				}
				return
			}
			if err != nil {
				t.Fatalf("Rot13(%q) unexpected error: %v", c.hexInput, err)
			}
			if got != c.expectHex {
				t.Errorf("Rot13(%q) == %q, want %q", c.hexInput, got, c.expectHex)
			}
		})
	}
}
