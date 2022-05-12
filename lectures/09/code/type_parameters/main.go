package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

type Tree[T interface{}] struct {
	left, right *Tree[T]
	value       T
}

func (t *Tree[T]) Lookup(x T) *Tree[T] {
	return nil
}

func main() {
	fmt.Println(min[int](1, 2), min[float32](2.0, 1.0), min[string]("foo", "bar"))
	fmt.Println(min(1, 2), min(2.0, 1.0), min("foo", "bar"))
	// min([]string{}, []string{}) // []string does not implement constraints.Ordered
	// min(1, 1.0) // default type float64 of 1.0 does not match inferred type int for T

	var stringTree Tree[string]
	stringTree.Lookup("foo")
	// stringTree.Lookup(1) // Cannot use '1' (type untyped int) as the type T (string)
}
