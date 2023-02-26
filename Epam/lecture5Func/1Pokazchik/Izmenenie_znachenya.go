package main

import "fmt"

func main() {
	var x = 100
	var p *int
	p = &x
	*p = 200 // установка нового значения в Х

	// Вывод адреса показчика
	fmt.Println("Value stored in variable p =", p)
	//Вывод измененного значения Р
	fmt.Println("Value stored in *p = ", *p)
	//Вывод измененного значения Х
	fmt.Println("Value stored in x = ", x)
}
