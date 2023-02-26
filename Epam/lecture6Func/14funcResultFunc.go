package main

import "fmt"

func add(x int, y int) int      { return x + y }
func subtract(x int, y int) int { return x - y }
func multiply(x int, y int) int { return x * y }

func selectFn(n int) func(int, int) int {
	if n == 1 {
		return add
	} else if n == 2 {
		return subtract
	} else {
		return multiply
	}
}

func selectFn2(n int) func(int, int) int {
	switch n {
	case 1:
		return add
	case 2:
		return subtract
	default:
		return multiply
	}
}

func main() {
	f := selectFn(1)
	fmt.Println(f(4, 3)) //7

	f = selectFn(3)
	fmt.Println(f(4, 3)) //12

	f = selectFn2(2)
	fmt.Print(f(4, 3)) //1
}
