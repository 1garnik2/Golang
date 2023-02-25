package main

import "fmt"

type employee struct {
	name      string
	age       int
	degnation string
	salary    int
}

func main() {
	var emploee1 *employee = &employee{"Rob", 43, "Developer", 3000}
	var emploee2 *employee = new(employee)

	fmt.Println("employee 1 :", *emploee1)
	fmt.Println("employee 2 :", *emploee2)

	// add values to employee2
	emploee2.name = "bobik"
	emploee2.age = 33
	emploee2.degnation = "trainee"
	emploee2.salary = 500
	fmt.Println("employee2 :", *emploee2)
	//
	fmt.Println(emploee2.name)
	fmt.Println((*emploee2).name)
}
