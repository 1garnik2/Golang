package main

import "fmt"

func main() {

	var students = map[string][]int{}

	students["Ihor"] = []int{94, 90}
	students["Iryna"] = []int{96, 90, 94}
	fmt.Println(students)
	fmt.Println(students["Iryna"])
	fmt.Println(students["Iryna"][1])

	var marks = map[[2]int]string{}
	marks[[2]int{90, 91}] = "vlad"
	marks[[2]int{91, 90}] = "Maryna"
	fmt.Println(marks)
	fmt.Println(marks[[2]int{91, 90}])
	fmt.Println(marks[[2]int{90, 91}])

}
