package main

import "fmt"

func main() {
	h, m, s := 0, 1, 1
	fmt.Println(Past(h, m, s))
}

func Past(h, m, s int) int {
	mil := 0 // your code here
	mil += s * 1000
	mil += m * 60000
	mil += h * 60 * 60000
	return mil
}
