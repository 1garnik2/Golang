package main

import "fmt"

func add(x int, y int) int      { return x + y }
func multiply(x int, y int) int { return x * y }

func display(massege string) {
	fmt.Println(massege)
}

func main() {
	//var f func(int, int) int = add
	// abo tak
	f := add
	fmt.Println(f(3, 4))
	f = multiply
	fmt.Println(f(4678, 8762))

	var s func(string) = display
	s("Hello")

}
