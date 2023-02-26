package main

import "fmt"

func main() {
	arr1 := [2]int{1, 2} // массив
	sum(arr1)
	fmt.Println(arr1)
	arr2 := []int{1, 2} //слайс
	sum2(arr2)
	fmt.Print(arr2)
}

// В значения в массиве
// меняються только в функции!
func sum(arr [2]int) {
	s := 0
	arr[0] = 100
	for _, x := range arr {
		s += x
	}
	fmt.Println(arr, "sum =", s)
}

// Слайс передаеться по адресу
// и значения в конечном слайсе изменяться!!!
func sum2(arr []int) {
	s := 0
	arr[0] = 100
	for _, x := range arr {
		s += x
	}
	fmt.Println(arr, "sum =", s)
}
