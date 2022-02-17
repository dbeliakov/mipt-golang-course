package sum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSum_Empty(t *testing.T) {
	require.Equal(t, 0, Sum(nil))
}

func TestSum_One(t *testing.T) {
	const num = 42
	require.Equal(t, num, Sum([]int{num}))
}

func TestSum_Couple(t *testing.T) {
	arr := []int{1, 2, 3, 10}
	require.Equal(t, 16, Sum(arr))
}

func TestSum_Negative(t *testing.T) {
	arr := []int{1, 2, 3, -10}
	require.Equal(t, -4, Sum(arr))
}
