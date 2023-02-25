package main

import "fmt"

func main() {
	println()
	println()
	println()
	println()
	println()
	println()
	println()
	println()
	println()
	for y := 0; y < 6; y++ {
		for x := 0; x < 7; x++ {
			if (y == 0 && x%3 != 0) || (y == 1 && x%3 == 0) ||
				(y-x == 2) || (y+x == 8) {
				print("* ")
			} else {
				print("  ")
			}
		}
		fmt.Println()

	}
	println()
	println("Don't worry! Everything will be fine!")
	println("           I Love You!")
	println()
	println()
	println()
	println()
	println()
}
