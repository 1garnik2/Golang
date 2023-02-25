package main

import "fmt"

func main() {
	sum(2.6, 4.4, 9.6, 7)
	sum(1, 2.1)
	sum(3, 5, 9)
}

func sum(n ...float64) {
	s := 0.0
	for _, x := range n {
		s += x
	}
	fmt.Println("sum =", s)
}
