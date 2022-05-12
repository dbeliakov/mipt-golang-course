package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Scale returns a copy of s with each element multiplied by c.
func Scale[S ~[]E, E constraints.Integer](s S, c E) S {
	r := make(S, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}

type Point []int32

func (p Point) String() string {
	return fmt.Sprintf("Point{%v}", []int32(p))
}

func main() {
	fmt.Println(Scale(Point{1, 2}, 2).String())
}
