package main

import "fmt"

func main() {
	fmt.Print(gcd(48, 64))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
