package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d", i)
		}
	}
	fmt.Println()
	for i := 1; i <= 10; i++ {
		for k := 1; k <= 3; k++ {
			if i == 3 {
				break
			}
			fmt.Printf("%d", i)
		}

	}
	fmt.Println()

Outerloop:
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 3; j++ {
			if i == 3 {
				break Outerloop
			}
			fmt.Printf("%d", i)
		}
	}
	fmt.Println()
}
