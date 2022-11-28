package main

import "fmt"

func main() {
	sum([]int{1, 3}...) // три точки позволяют работать со слайсами
	sum([]int{3, 5, 9}...)
	////// abo tac
	var nums = []int{2, 5, 9}
	sum(nums...)

}

// непоределенное кол-во параметров
func sum(n ...int) {
	s := 0
	for _, x := range n {
		s += x
	}
	fmt.Println("sum =", s)
}
