package main

import "fmt"

func main() {
	var max int

	var a = [3][4]int{{1, 3, 74, 3}, {2, 73, 7, 9}, {54, 3, 5, 10}}

	max = a[0][0]

	for i := 0; i <= 2; i++ {
		for j := 0; j <= 3; j++ {
			if a[i][j] > max {
				max = a[i][j]

			}
		}
	}

	fmt.Printf("max = %d", max)
	fmt.Println()
}
