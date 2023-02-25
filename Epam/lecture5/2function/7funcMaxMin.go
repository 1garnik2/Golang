package main

import "fmt"

func main() {
	var x int
	var y int
	fmt.Scanf("%d", &x)
	fmt.Scanf("%d", &y)
	fmt.Println(max(x, y))
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
