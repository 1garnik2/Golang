// З перших  5 елементів масиву створити зріз.
//
//	Кожний елемент збільшити на максимум отриманого зрізу.
//
// Вивести отриманий зріз
package main

import "fmt"

func main() {
	arr := [...]int{48, 96, 86, 68, 57, 82, 63, 70}
	var slice []int
	slice = arr[:5]
	index := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] > slice[index] {
			index = i
		}
	}
	s := slice[index]
	for i := 0; i < len(slice); i++ {
		slice[i] += s
	}

	fmt.Println(slice)
}
