package main

import "fmt"

//создадим интерфейс
type animal interface {
	makeSound()
}

//===========================

// сдруктуры для интерфейса
type cat struct{}
type dog struct{}

//============================

// методы для интефейса
func (c *cat) makeSound() {
	fmt.Println("Myaw")
}
func (d *dog) makeSound() {
	fmt.Println("Woof")
}

//==========================
func main() {
	// две переменные для структур
	var c, d animal = &cat{}, &dog{}
	//=========================
	// Вызов метода makeSound() для структур
	c.makeSound()
	d.makeSound()
}
