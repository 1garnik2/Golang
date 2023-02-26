package main

import "fmt"

func main() {
	p1 := createPointer(7)
	fmt.Println("Адрес p1:", p1, "Значение *p1", *p1)
	p2 := createPointer(10)
	fmt.Println("Адрес p2:", p2, "Значение *p2", *p2)
	p3 := createPointer(28)
	fmt.Println("Адрес p3:", p3, "Значение *p3", *p3)

	x := 15
	//предача переменной в фунцию
	p4 := createPointer(x)
	fmt.Println("p4:", p4, "*p4", *p4)
}

//функция которая возвращает показчик
func createPointer(x int) *int {
	p := new(int) //создание показчика на целое число
	// abo tac p:=&x
	*p = x // значение инт которое передаеться
	return p
}
