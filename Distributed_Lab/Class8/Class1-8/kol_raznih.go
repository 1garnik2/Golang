package main

import (
	"fmt"
)

func main() {
	arr := [...]int{4, 3, 1, 5, 3, 14, 3, 4, 7, 10, 11, 12, 4, 3, 3, 56, 3, 12, 122, 3, 2, 3}
	var count int
	var temp int
	for i := 0; i < (len(arr) - 1); i++ {
		for j := 0; j < (len(arr)-1)-i; j++ {
			if arr[j] > arr[j+1] {
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}

	for i := 0; i < (len(arr) - 1); i++ {
		if arr[i] != arr[i+1] {
			count++
		}
	}
	fmt.Println(count + 1)
}
