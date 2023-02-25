// Построить слайс из перестановки
// Копирует слайс и значения первого ставит индексом второго
// Слайс не всегда копируеться правильно
package main

import "fmt"

func main() {
	slice := []int{0, 2, 1, 5, 3, 4}
	fmt.Println(buildArray(slice))
}

func buildArray(nums []int) []int {
	sol := make([]int, len(nums))
	for i, num := range nums {
		sol[i] = nums[num]
	}
	return sol
}
