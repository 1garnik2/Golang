package main

import (
	"fmt"
	"strings"
)

func main() {
	arr2 := [][]int{
		0: {0: 0, 1: 1, 2: 2, 3: 3, 4: 4},
		1: {0: 10, 1: 11, 2: 12, 3: 13, 4: 14},
		2: {0: 20, 1: 21, 2: 22, 3: 23, 4: 24},
		3: {0: 30, 1: 31, 2: 32, 3: 33, 4: 34},
	}
	fmt.Println(ToCsvText(arr2))
}

func ToCsvText(array [][]int) string {
	array2 := []string{}
	for _, a := range array {
		k := []string{}
		for _, n := range a {
			k = append(k, fmt.Sprint(n))
		}
		array2 = append(array2, strings.Join(k, ","))
	}
	return strings.Join(array2, "\n")
}
