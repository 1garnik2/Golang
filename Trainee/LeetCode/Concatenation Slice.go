// Копирует слайс два раза
// Возвращает увеличенный в два раза слайс
package main

import "fmt"

func main() {
	slice := []int{1, 3, 2, 1}
	fmt.Println(getConcatenation(slice))
}
func getConcatenation(nums []int) []int {
	//nums = append(nums, nums...) // увеличивает слайс в двое
	return append(nums, nums...)
}
