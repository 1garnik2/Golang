package main

import "fmt"

type Operation interface {
	sum() int
	mult() int
	devishion() float64
	substruct() int
}

func (n Numbers) sum() int {
	return (n.num1 + n.num2)
}

func (n Numbers) mult() int {
	return n.num1 * n.num2
}

func (n Numbers) devishion() float64 {
	return float64(n.num1) / float64(n.num2)
}

func (n Numbers) substruct() int {
	return n.num1 - n.num2
}

type Numbers struct {
	num1 int
	num2 int
}

func main() {
	var i Operation
	numbers := Numbers{3, 5}
	i = numbers
	//fmt.Println(numb.sum())
	fmt.Printf("%T\n", i)
	fmt.Printf("сумм %d \n", i.sum())
	fmt.Printf("умножение %d \n", i.mult())
	fmt.Printf("деление  %.4f\n", i.devishion())
	fmt.Printf("вычитание %d\n", i.substruct())

}
