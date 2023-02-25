package main

import . "fmt"

func main() {
	var n int
	Scan(&n)
	for d := 1; d <= n; d <<= 1 {
		Print(d, " ")
	}
}
