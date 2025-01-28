package rle

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRLECompress(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty",
			input:    "",
			expected: "",
		},
		{
			name:     "one repeating symbol",
			input:    "aaaaaaaaaaaaa",
			expected: "a13",
		},
		{
			name:     "non-repeating symbols",
			input:    "some string",
			expected: "s1o1m1e1 1s1t1r1i1n1g1",
		},
		{
			name:     "mixed",
			input:    "TTTATTAAAAC",
			expected: "T3A1T2A4C1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := RLECompress(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func BenchmarkRLECompress(b *testing.B) {
	input := strings.Repeat("GCTAGTTATTGGGG", 100)

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = RLECompress(input)
	}
}
