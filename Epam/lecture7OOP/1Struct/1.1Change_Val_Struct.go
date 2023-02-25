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
	fmt.Println(employee)
	// Изменение информации в структуре
	employee.age = 26
	employee.salary = 3000
	fmt.Println(employee.age)
	fmt.Println(employee.salary)
	fmt.Println(employee)
}
