package main

import "fmt"

type worker string

func (w worker) work() {
	fmt.Printf("%s is working\n", w)
}

type myType interface {
	work()
}

func DoWork[T myType](things []T) {
	for _, v := range things {
		v.work()
	}
}

func main() {
	var a, b, c worker
	a = "A"
	b = "B"
	c = "C"
	DoWork([]worker{a, b, c})
}
