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
		{name: "no args", strs: nil, expected: ""},
		{name: "two empty strings", strs: []string{"", ""}, expected: ""},
		{name: "one empty string", strs: []string{"", "Hello!"}, expected: ""},
		{name: "basic", strs: []string{"Hello World!", "Hello, World!"}, expected: "Hello"},
		{name: "equal strings", strs: []string{"abacaba", "abacaba"}, expected: "abacaba"},
		{name: "one is a prefix of another", strs: []string{"Hello, World!", "Hello"}, expected: "Hello"},
		{name: "no common prefix", strs: []string{"Hello", "hello"}, expected: ""},
		{name: "emoji", strs: []string{"😀", "😁"}, expected: ""},
		{name: "single string", strs: []string{"alone"}, expected: "alone"},
		{name: "three strings", strs: []string{"flower", "flow", "flight"}, expected: "fl"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := LongestCommonPrefix(tc.strs...)
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
