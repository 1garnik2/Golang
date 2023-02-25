package main

import "fmt"

func main() {
	//при изменении массива - измениться и слайс
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)
	// копия слайса
	getslice := slice1[:]
	fmt.Println(getslice)

}
