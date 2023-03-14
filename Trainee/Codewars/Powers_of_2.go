// слайс из 2 в степени n. При значении n=4 вывод {1,2,4,8,16}
package main

import (
	"fmt"
	"math"
)

func main() {
	n := 4
	fmt.Println(PowersOfTwo(n))
}

func PowersOfTwo(n int) []uint64 {
	result := make([]uint64, n+1)
	for i := 0; i <= n; i++ {
		result[i] = uint64(int(math.Pow(2, float64(i))))
	}
	return result
}
