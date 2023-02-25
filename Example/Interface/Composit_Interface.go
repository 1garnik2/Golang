package main

import "fmt"

//Интерфейс из интерфейсов
type animal interface {
	walker
	runner
}

type bird interface {
	walker
	flier
}

//интерфейс с методом walk
type walker interface {
	walk()
}

type runner interface {
	run()
}

type flier interface {
	fly()
}

//структура  cat
type cat struct{}
type eagle struct{}

// методы для структуры cat
func (c *cat) walk() {
	fmt.Println("Cat is walking")
}
func (c *cat) run() {
	fmt.Println("Cat is ranning")
}
func (e *eagle) walk() {
	fmt.Println("Eagle is walking")
}
func (e *eagle) fly() {
	fmt.Println("Eagle is flying")
}
func walk(w walker) {
	w.walk()
}
func main() {
	//переменная с типом интерфейс равна адресу структуры run
	var c animal = &cat{}
	var e bird = &eagle{}
	walk(c)
	walk(e)
	c.run()
	c.walk()
}
