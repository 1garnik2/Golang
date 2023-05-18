package main

import "fmt"

type DataStruct struct {
	param1 string
	param2 bool
	param3 *int
}

func generator[T any](length int) []T {
	res := make([]T, length)
	for i := 0; i < length; i++ {
		res[i] = *new(T)
	}
	return res
}
func main() {
	res := generator[string](10)
	fmt.Println(res)
	fmt.Println(generator[DataStruct](10))
	fmt.Println(generator[int](10))
}
