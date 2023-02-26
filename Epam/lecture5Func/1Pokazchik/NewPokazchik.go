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

	//Создание нового пустого показчика
	var p3 = new(int)
	//Вывод адреса нового показчика
	fmt.Println("Value stired in variable p3 = ", p3)
	//Вывод значения нового показчика
	fmt.Println("Value stored in *p3 = ", *p3)

}
