/*
Учитывая массив nums,

	 для каждого nums[i] узнайте,
	  сколько чисел в массиве меньше его.
	   То есть для каждого nums[i]
	    вы должны подсчитать количество
		 допустимых j таких,
	 что j != i и nums[j] < nums[i].
*/
package main

import "fmt"

func main() {
	arr := [4]int{6, 5, 4, 8}
	sl := arr[:]
	fmt.Println(smallerNumbersThanCurrent(sl))
}
func smallerNumbersThanCurrent(nums []int) []int {
	count := 0
	sol := make([]int, 0)
	for k, _ := range nums {
		count = 0
		for i := 0; i < len(nums); i++ {
			if nums[k] > nums[i] {
				count += 1
			}
		}

		sol = append(sol, count)
	}
	return sol
}
