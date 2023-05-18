package main

import "fmt"

func doSomething(r int, t int, callback func(int, int) int, s string) int {
	fmt.Println(s)
	return callback(r, t)
}

func main() {
	a, b := 2, 6
	result := doSomething(a, b, back, "multiplay")
	fmt.Printf("Multiply: %d\n", result)
	result = doSomething(a, b, sum, "sum")
	fmt.Printf("Sum: %d", result)
}

func back(n1, n2 int) int {
	return n1 * n2
}

func sum(n1, n2 int) int {
	return n1 + n2
}
