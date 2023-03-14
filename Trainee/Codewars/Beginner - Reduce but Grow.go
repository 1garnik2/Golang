package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(Grow(arr))
}
func Grow(arr []int) int {
	grow := 1
	for _, v := range arr {
		grow = grow * v
	}
	return grow
}
