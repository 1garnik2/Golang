package main

import "fmt"

func main() {
	type employee struct {
		name         string
		age          int
		designatoins string
		salary       int
	}
	emploee1 := employee{
		name:         "Rob Pyke",
		age:          25,
		designatoins: "Developer",
		salary:       2000,
	}
	emploee2 := employee{
		name:         "Ken Thompson",
		age:          23,
		designatoins: "Tester",
		salary:       1800,
	}
	fmt.Println(emploee1)
	fmt.Println(emploee2)
}
