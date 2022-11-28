package main

import "fmt"

func main() {
	sum(2, 4, 9, 7)
	sum(1, 2)
	sum(3, 5, 9)
}

func sum(n ...int) {
	s := 0
	for _, x := range n {
		s += x
	}
	fmt.Println("sum =", s)
}
