package main

import "fmt"

func main() {
	var a, b, x, o int32
	fmt.Scan(&a)
	fmt.Scan(&b)
	a *= 10
	x = (a / b)
	o = (a % b)
	if o > 0 {
		fmt.Println(x + 1)
	} else {
		fmt.Println(x)
	}

}
