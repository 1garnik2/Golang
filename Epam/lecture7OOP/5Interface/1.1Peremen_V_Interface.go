package main

import "fmt"

// интерфейс с методом
type Mover interface {
	move()
}

//структуры
type Human struct{}

type Car struct{}

type Plane struct{}

type Ship struct{}

// методы для структур с имплементацией метода move
func (h Human) move() {
	fmt.Println("The Human is going")
}

func (c Car) move() {
	fmt.Println("The Car is driving")
}
func (s Ship) move() {
	fmt.Println("The Ship is sailing")
}
func (p Plane) move() {
	fmt.Println("The Plane is flying")
}

//функция получает переменную типа интерфейс
func drive(m Mover) {
	m.move()
}
func main() {
	// переменные структуры
	var human = Human{}
	var car = Car{}
	var plane = Plane{}
	var ship = Ship{}
	//можно присваивать переменной
	var x Mover
	x = car
	x.move()

	x = human
	x.move()

	x = plane
	x.move()

	x = ship
	x.move()
}
