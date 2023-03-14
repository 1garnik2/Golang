package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, -3, 4, 5}

	fmt.Println(Invert(arr))
}

func Invert(arr []int) []int {
	result := make([]int, 0)
	for _, v := range arr {
		result = append(result, -v)
	}
	return result
}
