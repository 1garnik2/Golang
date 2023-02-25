package main

import "fmt"

type rabotnyk struct {
	name        string
	age         int
	designation string
	salary      int
}

//Создание метода setAge для структуры rabotnyk
func (metod *rabotnyk) setAge(newAge int) {
	(*metod).age = newAge
}
func (met rabotnyk) getAge() int {
	return met.age
}
func main() {
	var emploee = rabotnyk{"Rob Pyke", 34, "Developer", 2000}
	fmt.Println("Age:", emploee.age)
	emploee.setAge(22)
	fmt.Println("Age:", emploee.getAge())
}
