package main

import "fmt"

type employee struct {
	name         string
	age          int
	designatoins string
	salary       int
}

func main() {
	//в одну строку
	emploee1 := employee{"Rob Pike", 25, "Developer", 2000}
	//Можно обьявлять не все поля
	emploee2 := employee{
		name:         "Ken Thompson",
		age:          23,
		designatoins: "Tester",
	}
	//пустая переменная
	emploee3 := employee{}
	// заполнение переменной
	emploee3.age = 19
	emploee3.name = "Robert Griesemer"
	emploee3.designatoins = "Tranee"
	emploee3.salary = 600

	fmt.Println(emploee3)
	fmt.Println(emploee1)
	fmt.Println(emploee2)
}
