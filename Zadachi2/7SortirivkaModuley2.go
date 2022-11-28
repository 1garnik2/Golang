package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var slice []int
	var n, item int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&item)
		slice = append(slice, item)
	}
	sort.Slice(slice, func(i, j int) bool {
		return (math.Abs(float64(slice[i])) < math.Abs(float64(slice[j]))) ||
			(math.Abs(float64(slice[i])) == math.Abs(float64(slice[j])) && slice[j] > 0)
	})
	for _, i := range slice {
		fmt.Println(i, " ")
	}
}
