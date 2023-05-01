package main

import "fmt"

func main() {
	f := 8
	fmt.Println(Summation(f))
}

func Summation(n int) int {
	sum := 0
	for i := n; i > 0; i-- {
		sum += i
	}
	return sum
}
