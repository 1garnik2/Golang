package main

import (
	"fmt"
)

func work(x int) int {
	return x * x
}

func main() {
	const N = 10
	cache := map[int]int{}

	var n int
	for i := 0; i < N; i++ {
		_, err := fmt.Scan(&n)
		fmt.Println(n)
		if err != nil {
			panic(fmt.Sprintf("invalid int value: %d", n))
		}

		// Если значения нет в словаре
		if _, ok := cache[n]; !ok {
			cache[n] = work(n)
		}
		n = cache[n]

		fmt.Printf("%d ", n)
	}
}
