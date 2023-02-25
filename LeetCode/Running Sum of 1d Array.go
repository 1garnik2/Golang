// Скалдывает значения слайса
// вход {1, 2, 3, 4}
// выход [1 3 6 10]
package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 3, 4}
	sl := arr[:]
	fmt.Println(runningSum(sl))
}
func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}
	return nums
}
