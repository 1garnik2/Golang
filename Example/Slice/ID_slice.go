package main

import "fmt"

func main() {
	sl := []int64{1, 2, 3, 2, 2, 1, 4, 8, 6, 4, 5}
	fmt.Println(UniqueUserIDs(sl))
}

func UniqueUserIDs(userIDs []int64) []int64 {
	id := []int64{}
	seen := make(map[int64]bool)

	for _, v := range userIDs {
		if !seen[v] {
			id = append(id, v)
			seen[v] = true
		}
	}

	return id
}
