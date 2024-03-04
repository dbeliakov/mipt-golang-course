package filter

import (
	"reflect"
	"testing"
)

func TestFilterInPlace(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		predicate func(int) bool
		want      []int
	}{
		{
			name: "filter_even_numbers",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			predicate: func(x int) bool {
				return x%2 == 0
			},
			want: []int{2, 4, 6, 8, 10},
		},
		{
			name:      "empty_slice",
			nums:      []int{},
			predicate: func(x int) bool { return x%2 == 0 },
			want:      []int{},
		},
		{
			name: "all_elements_fail_predicate",
			nums: []int{1, 3, 5, 7, 9},
			predicate: func(x int) bool {
				return x%2 == 0
			},
			want: []int{},
		},
		{
			name: "all_elements_pass_predicate",
			nums: []int{2, 4, 6, 8},
			predicate: func(x int) bool {
				return x%2 == 0
			},
			want: []int{2, 4, 6, 8},
		},
		{
			name: "filter_positive_numbers",
			nums: []int{-3, -2, -1, 0, 1, 2, 3},
			predicate: func(x int) bool {
				return x > 0
			},
			want: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := FilterInPlace(tt.nums, tt.predicate)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("FilterInPlace() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
