package main

import "fmt"

type SomeInterface interface {
	SomeMethod(in string) (out string)
}

type SomeStruct struct {
	prefix string
}

func (s SomeStruct) SomeMethod(in string) (out string) {
	return s.prefix + in
}

func main() {
	var iv SomeInterface = SomeStruct{"I say: "}

	fmt.Printf("value %v\n", iv)
	fmt.Printf("pointer %v\n", &iv)
	fmt.Printf("method %q\n", iv.SomeMethod("test!"))
}
