package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a > 10 {
		goto L1
	}
	a *= 2
L1:
	a += 20
	fmt.Println(a)
}
