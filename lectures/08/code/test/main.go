package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f Foo

	val := reflect.ValueOf(&f)

	val.Elem().Field(0).SetInt(42)
	fmt.Println(f)

}

type Foo struct {
	I int
}
