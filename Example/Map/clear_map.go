package main

import "fmt"

func main() {
	//
	var students = make(map[string]int)
	students["Roman"] = 95
	students["Maryna"] = 93
	fmt.Println(students)
	//
	studentsList := make(map[string]int)
	studentsList["Oleksandr"] = 96
	studentsList["Alona"] = 94
	fmt.Println(studentsList)
	//
	for k := range students {
		delete(students, k)
	}
	fmt.Println(students)
	//
	studentsList = make(map[string]int)
	fmt.Println(studentsList)

}
