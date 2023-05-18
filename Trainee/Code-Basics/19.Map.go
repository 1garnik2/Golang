package main

import "fmt"

func main() {
	slice := []int64{3, 2, 5, 0, 3, 5, 4, 6, 4, 7, 5, 7, 6, 8, 9, 6, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(UniqueUserIDs(slice))
}

// BEGIN

// UniqueUserIDs removes duplicates from the userIDs slice saving the IDs order.
func UniqueUserIDs(userIDs []int64) []int64 {
	// пустая структура struct{} — это тип данных, который занимает 0 байPrintln()
	// используется, когда нужно проверять в мапе только наличие ключа
	processed := make(map[int64]struct{})

	uniqUserIDs := make([]int64, 0)
	for _, uid := range userIDs {
		if _, ok := processed[uid]; ok {
			continue
		}

		uniqUserIDs = append(uniqUserIDs, uid)
		processed[uid] = struct{}{}
	}
	fmt.Println(processed)
	return uniqUserIDs
}
