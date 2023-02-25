package main

import "fmt"

type employee struct {
	name         string
	age          int
	designatoins string
	salary       int
}

func main() {
	//в одну строку
	emploee1 := employee{"Rob Pike", 25, "Developer", 2000}
	//Copy
	var emploee2 employee = emploee1
	emploee2.age = 32

	fmt.Println(emploee2)
	fmt.Println(emploee1)
}
