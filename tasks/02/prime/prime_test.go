package prime_test

import (
	"github.com/dbeliakov/mipt-golang-course/tasks/02/prime"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNumbersN(t *testing.T) {
	for n := 2; n < 100; n++ {
		var nums = prime.Numbers(n)
		require.LessOrEqual(t, len(nums), len(primes))
		for i := range nums {
			require.Equal(t, primes[i], nums[i])
		}
		if len(nums) < len(primes) {
			require.LessOrEqual(t, n, primes[len(nums)])
		}
	}
}

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
