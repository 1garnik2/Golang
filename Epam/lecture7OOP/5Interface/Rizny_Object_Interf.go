package main

import "fmt"

//
type Swimer interface {
	swim()
}
type Flying interface {
	fly()
}

//
type Plane struct{}
type Ship struct{}
type Duck struct{}

//

func (p Plane) fly() {
	fmt.Println("The Plane is flying")
}
func (s Ship) swim() {
	fmt.Println("The Ship is sailing")
}

//////////////////////////////////////
func (d Duck) fly() {
	fmt.Println("The Duck is flying")
}
func (d Duck) swim() {
	fmt.Println("The Duck is sailing")
}

//////////////////////////////////////
func main() {
	var duck = Duck{}
	var plane = Plane{}
	var ship = Ship{}

	slice := []Flying{plane, duck}
	for _, i := range slice {
		i.fly()
	}

	slice1 := []Swimer{duck, ship}
	for _, i := range slice1 {
		i.swim()
	}

}
