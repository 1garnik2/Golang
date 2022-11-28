package main

import "fmt"

func main() {
	var a, b, n, s, g, k int32
	fmt.Scan(&a, &b, &n)
	a = a * 100
	s = (a + b)
	s = (s * n)
	g = s / 100
	k = s % 100
	fmt.Println(g, k)
}
