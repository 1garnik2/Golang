package main

import "fmt"

func main() {
	hello("World")
	hello("EPAM")
}
func hello(s string) {
	fmt.Println("Hello", s)
}
