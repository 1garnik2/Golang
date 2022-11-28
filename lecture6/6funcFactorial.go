package main

import "fmt"

func main() {
	a := 5

	fmt.Println(factorial(a))
	fmt.Println(factorial(31))
}
func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}
