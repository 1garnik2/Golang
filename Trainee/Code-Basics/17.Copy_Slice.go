package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(IntsCopy(nums, 3))
}

func IntsCopy(src []int, maxLen int) []int {
	if maxLen <= 0 {

		return []int{}
	}
	if maxLen > len(src) {
		maxLen = len(src)
	}
	sl := make([]int, maxLen)
	copy(sl, src)
	return sl
}
