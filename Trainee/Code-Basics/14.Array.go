package main

import "fmt"

func SafeWrite(nums [5]int, i, val int) [5]int {
	if i < len(nums) && i >= 0 {
		nums[i] = val
	}

	return nums
}
func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(SafeWrite(arr, 0, 35))
}
