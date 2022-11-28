package main

import "fmt"

func main() {

	sum2([]int{1, 3}) // без точек работать со слайсами
	sum2([]int{3, 5, 9})
	//abo tac
	var nums = []int{2, 4, 67}
	sum2(nums)

}

func sum2(arr []int) {
	s := 0
	for _, x := range arr {
		s += x
	}
	fmt.Println("sum2 =", s)
}
