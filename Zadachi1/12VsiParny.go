package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scanf("%d", &i)
	x := 2
	for x <= i {
		fmt.Printf("%d", &x)
		x += 2
	}
}
