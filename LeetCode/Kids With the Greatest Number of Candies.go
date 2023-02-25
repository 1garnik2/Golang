/*
Есть n детей с конфетами.
Вам дан целочисленный массив конфет,
где каждая конфета[i] представляет

	количество конфет, которое есть у i-го ребенка,
	 и целое число extraCandies,
	  обозначающее количество дополнительных конфет,

которые у вас есть.
*/
package main

import "fmt"

func main() {
	arr := []int{2, 3, 5, 1, 3}
	slice := arr
	fmt.Println(kidsWithCandies(slice, 3))
}
func kidsWithCandies(candies []int, extraCandies int) []bool {

	max := 0
	candibool := make([]bool, len(candies))
	for _, k := range candies {
		if k > max {
			max = k
		}
	}
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies < max {
			candibool[i] = false
		} else {
			candibool[i] = true
		}
	}
	return candibool
}
