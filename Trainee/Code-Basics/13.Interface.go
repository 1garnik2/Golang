package main

import "fmt"

type IVoiceable interface {
	Voice() string
}

type Cat struct{}
type Cow struct{}
type Dog struct{}

func (c Cat) Voice() string {
	return "Мяу"
}
func (cw Cow) Voice() string {
	return "Мууу"
}
func (d Dog) Voice() string {
	return "Гав"
}
func main() {
	cat := Cat{}
	cow := Cow{}
	dog := Dog{}
	fmt.Println(cat.Voice())
	fmt.Println(cow.Voice())
	fmt.Println(dog.Voice())
}
