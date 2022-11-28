package main

import "fmt"

func main() {

	var students = map[string]int{}

	fmt.Println(students)
	fmt.Printf("%Tn", students)
	fmt.Println(len(students))

	students["Ihor"] = 94
	students["Iryna"] = 92
	fmt.Println(students)
	fmt.Println(len(students))
}
