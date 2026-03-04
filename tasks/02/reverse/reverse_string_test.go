package reverse

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: ""},
		{name: "one symbol", input: "a", expected: "a"},
		{name: "basic string", input: "Hello, World!", expected: ",olleH !dlroW"},
		{name: "russian string", input: "Строка на русском", expected: "акортС ан мокссур"},
		{name: "sequences", input: "\r\n\t\b", expected: "\r\n\t\b"},
		{name: "simple hieroglyphs", input: "Hello, 世界", expected: ",olleH 界世"},
		{name: "complex hieroglyph", input: "뢴", expected: "ᆫᅬᄅ"},
		{name: "emoji", input: "🙂🙂", expected: "🙂🙂"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := ReverseString(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func BenchmarkReverseString(b *testing.B) {
	input := strings.Repeat("世界 🙂🙂 \rПривет, World!", 100)

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = ReverseString(input)
	}
}
