// Возвращает индексы массива сумма
// которых равна переданному числу в функции
package main

import "fmt"

func main() {
	var arr = [4]int{2, 5, 5, 11}
	fmt.Println(twoSum(arr, 10))
}

func twoSum(nums [4]int, target int) []int {
	slice := make([]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target && i != j {
				slice = append(slice, i)
				slice = append(slice, j)
				return slice
			}
		}
	}
	return slice
}
