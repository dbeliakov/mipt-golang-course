package main

import "fmt"

type UInt interface {
	uint | uint8 | uint16 | uint32 | uint64
}

func MaxUInt[T UInt](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func main() {
	// var x UInt // interface contains type constraints
	fmt.Println(MaxUInt(uint(1), 2))
	// MaxUInt("a", "b") // string does not implement UInt
}
