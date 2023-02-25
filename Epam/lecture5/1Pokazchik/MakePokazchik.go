package main

import "fmt"

func main() {
	//переменная
	var x = 100
	// установка показчика
	var p *int

	p = &x
	//вывод значения переменной
	fmt.Println("Value stored in x = ", x)
	//вывод адреса переменной
	fmt.Println("Adress of x =", &x)
	//вывод адреса переменной
	fmt.Println("value stored in variable = ", p)
	//вывод значения на которое ссылаеться показчик
	fmt.Println("Value stored in x(*p) = ", *p)
}
