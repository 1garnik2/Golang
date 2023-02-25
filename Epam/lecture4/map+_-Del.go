package main

import "fmt"

func main() {
	var students = map[string]int{"Oleksandr": 95, "Alona": 100, "Roman": 90, "Maryna": 92}
	fmt.Println(students)

	fmt.Println("Roman is rater", students["Roman"])
	fmt.Println("Olesandr is rated", students["Oleksandr"])

	//перезапись "оценки"
	students["Roman"] = 91
	fmt.Println("Roman is rater", students["Roman"])

	//Добавление ключа "студента с оценкой"

	students["Sergii"] = 96
	fmt.Println(students)

	// Удалекие ключа "студента"
	delete(students, "Oleksandr")
	fmt.Println(students)

	// Удаление map
	for k := range students {
		delete(students, k)
	}
	fmt.Println(students)
	// Удаление map
	studentsList := make(map[string]int)

	fmt.Println(studentsList)
}
