package main

import "fmt"

func main() {
	//
	var students = map[string]int{"Oleksandr": 95,
		"Alona": 100, "Roman": 90, "Maryna": 92}
	fmt.Println(students)
	//вывод значения ключа
	fmt.Println("Roman is rated", students["Roman"])
	//изменение значения ключа
	students["Roman"] = 94
	fmt.Println("Roman is rated", students["Roman"])

}
