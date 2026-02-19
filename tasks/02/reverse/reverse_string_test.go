package reverse

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseEachWord(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: ""},
		{name: "one word", input: "Hello", expected: "olleH"},
		{name: "two words", input: "Hello World", expected: "olleH dlroW"},
		{name: "russian words", input: "Привет мир", expected: "тевирП рим"},
		{name: "three words", input: "one two three", expected: "eno owt eerht"},
		{name: "hieroglyphs", input: "Hello 世界", expected: "olleH 界世"},
		{name: "single char words", input: "a b c", expected: "a b c"},
		{name: "mixed scripts", input: "Ура Go", expected: "арУ oG"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := ReverseEachWord(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func BenchmarkReverseEachWord(b *testing.B) {
	input := strings.Repeat("世界 Привет World Go", 100)

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = ReverseEachWord(input)
	}
}
