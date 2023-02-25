package main

import (
	"fmt"
)

func main() {
	var arr [100]int
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	// минимальное значение
	index := 0
	for i := 0; i < n; i++ {
		if arr[i] < arr[index] {
			index = i
			fmt.Println(arr[i])
		}

	}
	fmt.Println(arr[index])
	h := 0
	for a := 0; a < n; a++ {

		if arr[a]%2 == 0 {
			for ; a < n; a++ {
				h += arr[a]

			}

		}

	}
	fmt.Println(h)
}
