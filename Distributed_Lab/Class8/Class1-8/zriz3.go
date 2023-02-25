// З останніх чотирьох елементів масиву створити зріз.
// Звести у квадрат елементи, які менше останнього елементу зрізу,
// вивести суму елементів цього зрізу
package main

import "fmt"

func main() {
	arr := [...]int{48, 96, 86, 68, 57, 82, 63, 70}
	var slice []int
	slice = arr[4:]
	sum := 0
	for i := 0; i < (len(slice) - 1); i++ {
		if slice[i] < slice[3] {
			slice[i] = slice[i] * slice[i]
		}
		sum += slice[i]
	}
	fmt.Println(slice[3] + sum)

}

//3249
//3969
