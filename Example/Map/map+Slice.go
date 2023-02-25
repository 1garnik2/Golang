package main

import "fmt"

func main() {
	//инициализация мап
	var students = map[string][]int{}
	students["Ihor"] = []int{94, 90}
	students["Iryna"] = []int{96, 90, 94}
	//показ оценок от имени
	fmt.Println(students)
	fmt.Println(students["Iryna"])
	fmt.Println(students["Iryna"][1])
	// инициализация мап2
	var marks = map[[2]int]string{}
	marks[[2]int{90, 91}] = "Vlad"
	marks[[2]int{91, 90}] = "Maryna"
	// показ имени от оценок
	fmt.Println(marks)
	fmt.Println(marks[[2]int{90, 91}])
	fmt.Println(marks[[2]int{91, 90}])

}
