package main

import "fmt"

func main() {
	students1 := map[string]int{"Oleksandr": 95, "Alona": 100, "Roman": 90, "Maryna": 92}
	students2 := map[string]int{"Alona": 100, "Roman": 90, "Serhii": 91, "Eduard": 99}
	fmt.Println(students1)
	fmt.Println(students2)

	for k, v := range students2 {
		students1[k] = v

	}
	fmt.Println(students1)
}
