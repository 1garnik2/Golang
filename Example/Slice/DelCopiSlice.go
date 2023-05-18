package main

import (
	"fmt"
)

func main() {
	slice := []int{3, 2, 2, 3, 45, 67, 87, 9, 8, 7, 65, 432, 456, 4, 5, 0, 3}
	fmt.Println(RemoveDuplicates(slice))
}

func RemoveDuplicates(slice []int) []int {
	unique := []int{}

	for _, v := range slice {
		found := false
		for _, u := range unique {
			if u == v {
				found = true
				break
			}
		}
		if !found {
			unique = append(unique, v)
		}
	}

	return unique
}
