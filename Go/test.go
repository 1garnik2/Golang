package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var a = 1
	var b = 0
	for i := 0; i <= n; i++ {
		a, b = b, a+b
	}
	fmt.Println(a)

}
