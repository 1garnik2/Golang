package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5}

	sl2 := []string{"a", "b", "c", "d", "e", "f", "g"}
	ReverseSlice(sl)
	ReverseSlice(sl2)

}
func ReverseSlice[T any](s []T) {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		last--
		first++
	}
	fmt.Println(s)
}
