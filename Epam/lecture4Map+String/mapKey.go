package main

import (
	"fmt"
	"sort"
)

func main() {
	students := map[string]int{"Oleksandr": 95, "Alona": 100, "Roman": 90, "Marina": 92}

	fmt.Println(students)
	fmt.Println()

	for key, element := range students {
		fmt.Println("Key", key, "=>", "Element:", element)
	}
	fmt.Println()

	for k := range students {
		fmt.Println(k, students[k])

	}
	fmt.Println()
	// обьявление слайса от map
	keys := make([]string, 0, len(students))

	for k := range students {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println(keys)

	for _, k := range keys {
		fmt.Println(k, students[k])
	}
	//создание слайса инт
	values := make([]int, 0, len(students))

	for _, v := range students {
		values = append(values, v)
	}
	// сортировка слайса инт
	sort.Ints(values)

	for _, v := range values {
		fmt.Println(v)
	}

}
