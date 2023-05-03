// Удаление по индеку слайса
package main

import "fmt"

func Remove(nums []int, i int) []int {
	if i < 0 || i > len(nums)-1 {
		return nums
	}
	slice := []int{}
	slice = append(slice, nums[:i]...)
	slice = append(slice, nums[i+1:]...)

	// nums[i] = nums[len(nums)-1]
	// fmt.Println(len(nums))
	// fmt.Println(nums[i])
	// fmt.Println(nums[:len(nums)-1])
	return slice
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(Remove(nums, 0))

}
