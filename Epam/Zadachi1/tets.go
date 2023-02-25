package main

import "fmt"

func main() {
	var max int

	var a = [3][4]int{{9, 5, 7, 6}, {6, 7, 7, 9}, {22, 8, 3, 10}}

	for i := 0; i <= 2; i++ {
		max = a[0][0]
		for j := 0; j <= 3; j++ {
			if a[i][j] < max {
				max = a[i][j]

			}
		}
		fmt.Print(max, " ")
	}

}
