package main

import "fmt"

func main() {
	x := 2
	y := 4
	sum1(x, y)
	a := sum2(x, y)
	fmt.Println(a)
	//fmt.Println(sum1(x, y)) Hельзя использовать т.к. не возвращает результат

	//Mожно использовать только когда функция ВОЗВРАЩАЕТ результат
	fmt.Println(sum2(x, y))
	//умножает х*у и складывает func
	sum1(sum2(3, 2), sum2(2, 6))
}

//без возврата результата
func sum1(x int, y int) {
	fmt.Println(x + y)
}

//возврат резутьтата
func sum2(x int, y int) int {
	return x * y
}
