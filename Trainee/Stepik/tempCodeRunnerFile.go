package main

import "fmt"

func main() {
	var n, s int
	fmt.Scan(&n)
	s = 1
	for s <= n {
		if s > n {
			break
		}
		fmt.Print(s, " ")
		s = s * 2
	}
}
