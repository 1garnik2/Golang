package main

import (
	"fmt"
	"sort"
)

func main() {

	intslice := []int{5, 2, 3, 10, 4, 7, 9, 1, 8, 6}
	strslice := []string{"Yuliia", "Serhii", "Alona", "Roman", "Victoriia", "Oleksandr"}

	//сортировка инт
	sort.Ints(intslice)
	fmt.Println(intslice)

	//сортирова стринг
	sort.Strings(strslice)
	fmt.Println(strslice)

	//бинарный поиск найти позицию єлемента инт
	i := sort.SearchInts(intslice, 6)
	fmt.Printf("found %d at index %d in %v\n", 6, i, intslice)

	//бинарный поиск найти позицию єлемента стринг
	s := sort.SearchStrings(strslice, "Roman")
	fmt.Printf("found %s at index %d in %v\n", "Roman", s, strslice)
	//компаратор от большего к меньшему инт
	sort.Slice(intslice, func(i, j int) bool {
		return intslice[i] > intslice[j]
	})
	fmt.Println(intslice)

	//компаратор от большего к меньшему последнего варианта!!!
	//сортировка сначала четных потом не четных
	sort.Slice(intslice, func(i, j int) bool {
		return intslice[i]%2 == 0 && intslice[j]%2 != 0
	})
	fmt.Println(intslice)

	//сортировка по длинне стринг имен
	//от коротких к длинным
	sort.Slice(strslice, func(i, j int) bool {
		return len(strslice[i]) < len(strslice[j])
	})
	fmt.Println(strslice)

}
