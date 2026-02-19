package lcp

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []struct {
		name     string
		strs     []string
		expected string
	}{
		{name: "empty slice", strs: []string{}, expected: ""},
		{name: "single string", strs: []string{"hello"}, expected: "hello"},
		{name: "two empty strings", strs: []string{"", ""}, expected: ""},
		{name: "one empty string", strs: []string{"", "Hello!"}, expected: ""},
		{name: "basic two", strs: []string{"Hello World!", "Hello, World!"}, expected: "Hello"},
		{name: "three strings", strs: []string{"Hello, World!", "Hello there", "Hello!"}, expected: "Hello"},
		{name: "equal strings", strs: []string{"abacaba", "abacaba", "abacaba"}, expected: "abacaba"},
		{name: "one is prefix of another", strs: []string{"Hello, World!", "Hello"}, expected: "Hello"},
		{name: "no common prefix", strs: []string{"Hello", "hello"}, expected: ""},
		{name: "emoji", strs: []string{"😀", "😁"}, expected: ""},
		{name: "four strings", strs: []string{"abc", "abd", "abe", "abz"}, expected: "ab"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := LongestCommonPrefix(tc.strs...)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func BenchmarkLongestCommonPrefix(b *testing.B) {
	strs := []string{
		strings.Repeat("a", 100) + "😀",
		strings.Repeat("a", 100) + "😁",
		strings.Repeat("a", 100) + "😂",
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = LongestCommonPrefix(strs...)
	}
}
