package main

import (
	"fmt"
	"sort"
)

func main() {

	intSlice := []int{5, 2, 1, 3, 10, 4, 7, 9, 8, 6}
	strSlice := []string{"Yuliia", "Serhii", "Alona", "Roman", "Victoriia", "Oleksandr"}

	sort.Ints(intSlice)
	fmt.Println(intSlice)

	sort.Strings(strSlice)
	fmt.Println(strSlice)

	// Бинарный поиск
	i := sort.SearchInts(intSlice, 6)
	fmt.Printf("found %d at index %d in %v\n", 6, i, intSlice)

	s := sort.SearchStrings(strSlice, "Roman")
	fmt.Printf("found %s at index %d in %v\n", "Roman", s, strSlice)
	//сортировка от большего к меньшему КОМПАРАТОР
	sort.Slice(intSlice, func(i, j int) bool {
		return intSlice[i] > intSlice[j]
	})
	fmt.Println(intSlice)

	//сортировка от большего к меньшему КОМПАРАТОР
	sort.Slice(intSlice, func(i, j int) bool {
		return intSlice[i]%2 == 0 && intSlice[j]%2 != 0
	})
	fmt.Println(intSlice)
	// сортирова строк
	sort.Slice(strSlice, func(i, j int) bool {
		return len(strSlice[i]) < len(strSlice[j])
	})
	fmt.Println(strSlice)

}
