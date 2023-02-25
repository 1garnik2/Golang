// Перетосовать слайс [x1,x2,...,xn,y1,y2,...,yn].
// вернуть [x1,y1,x2,y2,...,xn,yn].

package main

import "fmt"

func main() {
	arr := [6]int{2, 5, 1, 3, 4, 7}
	slice := arr[:]
	fmt.Println(shuffle(slice, 3))
}
func shuffle(nums []int, n int) []int {
	result := []int{}

	for i := 0; i < n; i++ {
		result = append(result, nums[i], nums[i+n])
	}

	return result
}
