package main

import (
	"fmt"
	"sort"
)

func main() {
	students := map[string]int{"Oleksandr": 95, "Alona": 100, "Roman": 90, "Maryna": 95}

	fmt.Println(students)

	//создание слайса от мап
	keys := make([]string, 0, len(students))
	for k := range students {
		keys = append(keys, k)
	}
	// сортировка по ключу и вывод слайса
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, students[k])
	}

	// создание слайса по значению
	values := make([]int, 0, len(students))

	for _, v := range students {
		values = append(values, v)
	}
	//сортировка по значению
	sort.Ints(values)
	// вывод отсортированного слайса в порядке возростания
	for _, v := range values {
		fmt.Println(v)
	}
}
