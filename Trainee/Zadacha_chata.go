package main

import "fmt"

type Person struct {
	name   string
	age    int
	Gender string
}

func main() {
	Slice := []Person{
		{"Mira", 27, "Women"},
		{"Nazar", 3, "Men"},
		{"Ihor", 38, "Men"},
	}
	for _, v := range Slice {
		fmt.Printf("Name : %s, Age : %d, Gender %s\n", v.name, v.age, v.Gender)
	}
}
