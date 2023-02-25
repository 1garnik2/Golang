package main

import "fmt"

func main() {
	var v, t, i, a int
	fmt.Scan(&v, &t)
	i = ((v * t) % 109)
	a = ((i + 109) % 109)
	fmt.Println(a)
}
