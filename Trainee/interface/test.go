package main

import (
	"fmt"
	"strings"
)

type Test interface {
	myMethod() string
	myInterface
}

type myInterface interface {
	myMethod() string
}

type mySlice []string

func (s mySlice) myMethod() string {
	return strings.Join(s, ", ")
}

func main() {
	var myVar myInterface = mySlice{"Hello", "world", "!"}
	fmt.Println(myVar.myMethod()) // Выводит: Hello, world, !
}
