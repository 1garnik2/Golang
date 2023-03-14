package main

import (
	"fmt"
)

func main() {
	groupCity := map[int][]string{
		10:   []string{"city1", "city2", "city3"}, // города с населением 10-99 тыс. человек
		100:  []string{"city4", "city5", "city6"}, // города с населением 100-999 тыс. человек
		1000: []string{"city7", "city8", "city9"}, // города с населением 1000 тыс. человек и более
	}

	cityPopulation := map[string]int{
		"city4": 101,
		"city5": 102,
		"city6": 103,
		"city1": 1,
		"city9": 1000,
	}
	for key, _ := range cityPopulation {
		isInMap := false
		for _, zn := range groupCity[100] {
			if key == zn {
				isInMap = true
			}
		}
		if isInMap == false {
			delete(cityPopulation, key)
		}
	}

	fmt.Print(cityPopulation)
}
