package genericutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFilter(t *testing.T) {
	t.Run("filter even numbers", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		even := Filter(nums, func(n int) bool { return n%2 == 0 })
		assert.Equal(t, []int{2, 4}, even)
	})

	t.Run("filter empty slice", func(t *testing.T) {
		var empty []string
		result := Filter(empty, func(s string) bool { return len(s) > 0 })
		assert.Empty(t, result)
	})
}

func TestGroupBy(t *testing.T) {
	type person struct {
		Name string
		Age  int
	}

	t.Run("group by age", func(t *testing.T) {
		people := []person{
			{"Alice", 25},
			{"Bob", 25},
			{"Charlie", 30},
		}
		groups := GroupBy(people, func(p person) int { return p.Age })
		expected := map[int][]person{
			25: {{"Alice", 25}, {"Bob", 25}},
			30: {{"Charlie", 30}},
		}
		assert.Equal(t, expected, groups)
	})

	t.Run("empty input", func(t *testing.T) {
		var empty []person
		groups := GroupBy(empty, func(p person) string { return p.Name })
		assert.Empty(t, groups)
	})
}

func TestMaxBy(t *testing.T) {
	t.Run("max by length", func(t *testing.T) {
		words := []string{"a", "aa", "aaa"}
		max := MaxBy(words, func(a, b string) bool { return len(a) < len(b) })
		assert.Equal(t, "aaa", max)
	})

	t.Run("empty slice returns zero value", func(t *testing.T) {
		var empty []int
		result := MaxBy(empty, func(a, b int) bool { return a < b })
		assert.Zero(t, result)
	})
}

func TestRepeat(t *testing.T) {
	t.Run("repeat value", func(t *testing.T) {
		result := Repeat("x", 3)
		assert.Equal(t, []string{"x", "x", "x"}, result)
	})

	t.Run("zero times", func(t *testing.T) {
		result := Repeat(1, 0)
		assert.Empty(t, result)
	})
}

func TestJSONParse(t *testing.T) {
	t.Run("parse valid JSON", func(t *testing.T) {
		type data struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		jsonStr := []byte(`{"name":"Alice","age":25}`)

		result, err := JSONParse[data](jsonStr)
		require.NoError(t, err)
		assert.Equal(t, data{"Alice", 25}, result)
	})

	t.Run("parse invalid JSON", func(t *testing.T) {
		_, err := JSONParse[int]([]byte("invalid"))
		require.Error(t, err)
	})
}

func TestDedup(t *testing.T) {
	t.Run("deduplicate integers", func(t *testing.T) {
		nums := []int{1, 2, 2, 3, 3, 3}
		unique := Dedup(nums)
		assert.ElementsMatch(t, []int{1, 2, 3}, unique)
	})

	t.Run("deduplicate strings", func(t *testing.T) {
		words := []string{"a", "b", "a", "c"}
		unique := Dedup(words)
		assert.ElementsMatch(t, []string{"a", "b", "c"}, unique)
	})
}
