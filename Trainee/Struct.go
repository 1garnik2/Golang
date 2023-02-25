package main

import "fmt"

type emploee struct {
	name     string
	age      int
	position string
	salary   int
}

func (e emploee) getName() string {
	return e.name
}
func (e emploee) getAge() {
	fmt.Println(e.age)
}
func main() {
	var rabotnik emploee
	rabotnik = emploee{"Bob", 23, "Trainee", 345}
	fmt.Println(rabotnik.position)
	rabotnik.getAge()
	rabotnik.name = "Rob"
	fmt.Print(rabotnik.getName())

}
