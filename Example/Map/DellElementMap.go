package main

import "fmt"

func main() {
	//
	var students = map[string]int{"Oleksandr": 95,
		"Alona": 100, "Roman": 90, "Maryna": 92}
	fmt.Println(students)
	//добавление элемента ключа и значения
	students["Serhii"] = 96
	fmt.Println(students)
	//удаление єлемента с ключом из map
	delete(students, "Oleksandr")
	fmt.Println(students)
}
