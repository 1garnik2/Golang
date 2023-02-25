package main

import "fmt"

func pr(prime int) bool {
	for i := 2; i < prime; i++ {
		if prime%i == 0 {
			return true
		}
	}
	return false
}
func main() {
	var prime int
	fmt.Scan(&prime)
	fmt.Print(pr(prime))
}
