package main

import "fmt"

func main() {
	//обьявляя переменную типом интерфейс
	// - можно присваивать ей любые значения
	var a interface{}
	a = 45
	fmt.Printf("%v,%T\n", a, a)
	a = "test"
	fmt.Printf("%v,%T\n", a, a)
	a = true
	fmt.Printf("%v,%T\n", a, a)
	a = 5.678
	fmt.Printf("%v,%T\n", a, a)
}
