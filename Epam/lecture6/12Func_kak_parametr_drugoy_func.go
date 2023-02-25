package main

import "fmt"

func add(x int, y int) int {
	return x + y

}
func multiply(x int, y int) int {
	return x * y
}

//Функция...10......25..  имя передаваемой фунции
func action(n1 int, n2 int, operation func(int, int) int) {
	result := operation(n1, n2)
	fmt.Println(result)
}
func main() {
	action(10, 25, add)
	action(5, 6, multiply)

}
