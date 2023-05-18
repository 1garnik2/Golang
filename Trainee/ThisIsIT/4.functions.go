package main

import "fmt"

func main() {

	first, second := 9, 4

	sumFunc := func(x, y int) int { return x + y } // переменная типа функция

	subtractFunc := func(x, y int) int { return x - y } // переменная типа функция

	fmt.Println(calculate(first, second, sumFunc))
	fmt.Println(calculate(first, second, subtractFunc))

	var multilaier func(x, y int) int // переменная типа функция

	multilaier = func(x, y int) int { return x * y } // инициализация переменной = функция
	fmt.Println(multilaier(first, second))

	divider := func(x, y int) float64 { return float64(x) / float64(y) } // переменная типа функция
	fmt.Println(divider(first, second))

}

// функция которая принимает х у + функцию и возвращает float64
func calculate(x, y int, action func(x, y int) int) float64 {
	return float64(action(x, y))
}
