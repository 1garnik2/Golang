package main

import "fmt"

type A struct {
	val  int64
	val2 int64
}

type B struct {
	val  float64
	val2 float64
}

type Str interface {
	getVal() interface{}
	getVal2() interface{}
}

func (a A) getVal() interface{} {
	return float64(a.val)
}
func (a A) getVal2() interface{} {
	return float64(a.val2)
}

func (b B) getVal() interface{} {
	return b.val
}
func (b B) getVal2() interface{} {
	return b.val2
}

func sumValues[T Str](Inter T) interface{} {

	return Inter.getVal().(float64) + Inter.getVal2().(float64)
}

func main() {

	a := A{val: 1, val2: 2}
	b := B{val: 1.1, val2: 2.2}

	fmt.Println(sumValues(a))
	fmt.Println(sumValues(b))
}
