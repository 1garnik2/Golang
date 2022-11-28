package main

import "fmt"

func main() {
	var arr = [10]int{1, 2, -3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(arr)
	index := 0
	for i := 0; i < 10; i++ {
		if arr[i] < arr[index] {
			index = i
		}
	}
	fmt.Println(arr[index])
}
