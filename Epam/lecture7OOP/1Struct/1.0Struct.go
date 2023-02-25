package main

import "fmt"

func main() {
	employee := struct {
		name         string
		age          int
		designatoins string
		salary       int
	}{
		name:         "Rob Pike",
		age:          25,
		designatoins: "Developer",
		salary:       2000,
	}
	// Вывод всей информации про структуру
	fmt.Println(employee)
	//Вывод полей структуры по отдельности
	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.designatoins)
	fmt.Println(employee.salary)
}
