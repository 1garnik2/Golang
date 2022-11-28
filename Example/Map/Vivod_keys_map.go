package main

import (
	"fmt"
	"sort"
)

func main() {
	students := map[string]int{"Oleksandr": 95, "Alona": 100, "Roman": 90, "Maryna": 95}

	fmt.Println(students)

	for key, element := range students {
		fmt.Println("Key", key, "=>", "Element:", element)
	}
	fmt.Println()

	for k := range students {
		fmt.Println(k, students[k])
	}
	fmt.Println()

	//
	keys := make([]string, 0, len(students))
	for k := range students {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, students[k])
	}
}
