package main

import "fmt"

type Foo struct{}

func (f Foo) String() string {
	return ""
}

var _ fmt.Stringer = Foo{}
var _ fmt.Stringer = &Foo{}

type Bar struct{}

func (b *Bar) String() string {
	return ""
}

// var _ fmt.Stringer = Bar{} // error: Type does not implement 'fmt.Stringer' as 'String' method has a pointer receiver
var _ fmt.Stringer = &Bar{}

// Nil interfaces

type Err struct{}

func (Err) Error() string {
	return ""
}

func foo() *Err {
	return nil
}

func bar() error {
	return foo()
}

func main() {
	err := bar()
	fmt.Printf("%v == nil? %v", err, err == nil)
}
