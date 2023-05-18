package main

import (
	"fmt"
	"strconv"
)

type DataStruct[T any] struct {
	ID   int64
	Data T
}

func collectData[T any](arr []DataStruct[T]) map[int64]T {
	res := make(map[int64]T)

	for i := 0; i < len(arr); i++ {
		res[arr[i].ID] = arr[i].Data
	}
	return res
}

func main() {
	data := make([]DataStruct[string], 8)

	for i := 0; i < len(data); i++ {
		data[i] = DataStruct[string]{
			ID:   int64(i),
			Data: strconv.Itoa(i * 34),
		}
	}
	fmt.Println(collectData(data))

}
