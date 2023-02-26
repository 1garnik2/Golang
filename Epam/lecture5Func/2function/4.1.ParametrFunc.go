package main

import "fmt"

func main() {
	hello("Hello", "Epam")
	hello("Hi", "World")
}

func hello(d, r string) {
	fmt.Println(d, r)
}
