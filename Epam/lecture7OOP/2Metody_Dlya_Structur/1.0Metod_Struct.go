package main

import "fmt"

type employee struct {
	name   string
	age    int
	design string
	salary int
}

// print all struct
func (met employee) aboutMe() {
	fmt.Println("Name:", met.name)
	fmt.Println("Age:", met.age)
	fmt.Println("Desingn:", met.design)
	fmt.Println("Salary:", met.salary)
}
func (met employee) getAge() int {
	return met.age
}
func main() {
	var emploee1 = employee{"Rob Pyke", 34, "Developer", 2000}
	fmt.Println(emploee1)
	emploee1.aboutMe()
	fmt.Println("Age:", emploee1.getAge())
}
