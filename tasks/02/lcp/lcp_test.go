package lcp

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []struct {
		name     string
		l        string
		r        string
		expected string
	}{
		{name: "two empty strings", l: "", r: "", expected: ""},
		{name: "one empty string", l: "", r: "Hello!", expected: ""},
		{name: "basic", l: "Hello World!", r: "Hello, World!", expected: "Hello"},
		{name: "equal strings", l: "abacaba", r: "abacaba", expected: "abacaba"},
		{name: "one is a prefix of another", l: "Hello, World!", r: "Hello", expected: "Hello"},
		{name: "no common prefix", l: "Hello", r: "hello", expected: ""},
		{name: "emoji", l: "😀", r: "😁", expected: ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := LongestCommonPrefix(tc.l, tc.r)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func BenchmarkLongestCommonPrefix(b *testing.B) {
	l, r := strings.Repeat("a", 100)+"😀", strings.Repeat("a", 100)+"😁"

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = LongestCommonPrefix(l, r)
	}
}
