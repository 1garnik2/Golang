package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int64{2, 1, 6, 5, 3, 4}
	fmt.Println(UniqueSortedUserIDs(nums))
}

func UniqueSortedUserIDs(userIDs []int64) []int64 {
	uniqueMap := make(map[int64]bool)
	uniqueIDs := userIDs[:0] // создаем слайс, указывающий на ту же память, что и userIDs

	for _, id := range userIDs {
		if !uniqueMap[id] {
			uniqueMap[id] = true
			uniqueIDs = append(uniqueIDs, id)
		}
	}

	sort.Slice(uniqueIDs, func(i, j int) bool {
		return uniqueIDs[i] < uniqueIDs[j]
	})

	return uniqueIDs
}
