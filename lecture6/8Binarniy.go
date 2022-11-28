package main

import "fmt"

func main() {
	printBinary(12)
	fmt.Println()

}
func printBinary(x int) {
	if x == 0 {
		return
	}
	printBinary(x / 2)
	fmt.Print(x % 2)
}
