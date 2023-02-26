package main

import "fmt"

// ввод 10 чисел для масива
func main() {
	var mas [10]int
	for i := 0; i < 10; i++ {
		fmt.Scan(&mas[i])
	}
	fmt.Println(mas)

	// сумма чисел масива
	sum := 0
	for i := 0; i < 10; i++ {
		sum += mas[i]
	}
	fmt.Println("sum=", sum)

	// максимальное значение
	//Обращение к элементу за индексом
	index := 0
	for i := 1; i < 10; i++ {
		if mas[i] > mas[index] {
			index = i
		}
	}
	fmt.Println(index) // вывод индекса максимального числа
	fmt.Println("max=", mas[index])

	// разворот массива
	var temp int
	for i := 0; i < 5; i++ {
		temp = mas[i]
		mas[i] = mas[9-i]
		mas[9-i] = temp
	}
	fmt.Println(mas)
}
