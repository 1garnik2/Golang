package main

import "fmt"

func main() {
	var x = 100

	var p *int
	p = &x
	fmt.Println("Value stored in x = ", x)
	//вывод адреса переменной
	fmt.Println("Adress of x =", &x)
	//вывод адреса переменной
	fmt.Println("value stored in variable = ", p)
	//вывод значения на которое ссылаеться показчик
	fmt.Println("Value stored in x(*p) = ", *p)
	//
	var p2 = &p
	//Вывод присвоеного адреса показчика Р2
	fmt.Println("Value stored in variable p2 =", p2)
	//Вывод адреса показчика Р через значение показчика Р2
	fmt.Println("Value stored in *p2 = ", *p2)
	//Вывод значения Х через показчика на показчик
	fmt.Println("value stored in **p2 = ", **p2)
}
