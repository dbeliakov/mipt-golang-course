package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x = 2
	v := reflect.ValueOf(x)
	fmt.Println(v.CanAddr())
	// v.Set(reflect.ValueOf(3)) // panic: reflect: reflect.Value.Set using unaddressable value
	fmt.Println("===")

	vp := reflect.ValueOf(&x).Elem()
	fmt.Println(vp.CanAddr())
	px := vp.Addr().Interface().(*int)
	*px = 3
	fmt.Println(x)
	fmt.Println("===")

	vp.Set(reflect.ValueOf(4))
	fmt.Println(x)
	fmt.Println("===")

	vp.SetInt(5)
	fmt.Println(x)
	fmt.Println("===")

	var i interface{}
	v = reflect.ValueOf(&i).Elem()
	fmt.Println(v.CanAddr())
	// v.SetInt(1) // panic: reflect: call of reflect.Value.SetInt on interface Value
	v.Set(reflect.ValueOf(1))
	fmt.Printf("%T: %v\n", i, i)
}
