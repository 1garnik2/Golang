package main

import "fmt"

func main() {

	var students = map[string]int{}

	fmt.Println(students)
	fmt.Printf("%T\n", students)
	fmt.Println(len(students))

	students["Ihor"] = 94
	students["Iryna"] = 92
	fmt.Println(students)
	fmt.Println(len(students))
}
