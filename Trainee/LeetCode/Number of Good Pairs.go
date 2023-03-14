// Количество одинаковых пар
package main

import "fmt"

func main() {
	array := [6]int{1, 2, 3, 1, 1, 3}
	sl := array[:]
	fmt.Println(numIdenticalPairs(sl))
}
func numIdenticalPairs(nums []int) int {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i] == nums[j] && i < j {
				count++
			}
		}
	}
	return count
}
