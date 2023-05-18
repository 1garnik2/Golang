// копирует слайс , добавляет знак !
// Исходный слайс остаеться без изменений
package main

import (
	"fmt"
)

func Map(strs []string, mapFunc func(s string) string) []string {
	mappedStrs := make([]string, len(strs))
	for i, s := range strs {
		mappedStrs[i] = mapFunc(s)
	}
	return mappedStrs
}

func main() {
	// пример использования функции Map
	inputStrs := []string{"hello", "world", "gopher"}
	mappedStrs := Map(inputStrs, func(s string) string {
		return s + "!"
	})
	fmt.Println(mappedStrs) // ["hello!", "world!", "gopher!"]
	// исходный слайс
	fmt.Println(inputStrs)
}
