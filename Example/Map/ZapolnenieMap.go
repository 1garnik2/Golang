package main

import "fmt"

func main() {
	var students = map[string]int{}
	fmt.Println(students)
	fmt.Printf("%T\n", students)
	fmt.Println(len(students))
	// добавление єлементов в пустой map
	students["Ihor"] = 94
	students["Iryna"] = 92
	fmt.Println(students)
	fmt.Println(len(students))

	//при добавлении одинкавых эелементов
	// в пустой map значение перезаписывается по последнему элементу
	students["Ihor"] = 94
	students["Ihor"] = 92
	fmt.Println(students)
	fmt.Println(len(students))
}
