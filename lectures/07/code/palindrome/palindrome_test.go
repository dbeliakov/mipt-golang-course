package palindrome

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	var testCases = []struct {
		string       string
		isPalindrome bool
		name         string
	}{
		{
			string:       "detartrated",
			isPalindrome: true,
		},
		{
			string:       "kayak",
			isPalindrome: true,
		},
		{
			string:       "palindrome",
			isPalindrome: false,
		},
		{
			string:       "été",
			isPalindrome: true,
			name:         "french",
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.isPalindrome, IsPalindrome(tc.string))
		})
	}
}

// func TestRandomPalindromes(t *testing.T) {
// 	seed := time.Now().UTC().UnixNano()
// 	t.Logf("Random seed = %v", seed)
// 	rnd := rand.New(rand.NewSource(seed))
// 	for i := 0; i < 1000; i++ {
// 		p := randomPalindrome(rnd)
// 		if !IsPalindrome(p) {
// 			t.Fail()
// 		}
// 	}
// }

func randomPalindrome(rnd *rand.Rand) string {
	n := rnd.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rnd.Intn(0x1000)) // '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}
