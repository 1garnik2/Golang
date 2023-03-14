// Ите максимальную сумму
// из массивов в двумерном массиве
package main

import "fmt"

func main() {
	var arr = [][]int{{1, 5}, {7, 3}, {3, 5}}

	fmt.Println(maximumWealth(arr))
}
func maximumWealth(accounts [][]int) int {
	sol := 0
	v := 0
	for k := 0; k < len(accounts); k++ {
		v = 0
		for _, i := range accounts[k] {
			v += i
			if v > sol {
				sol = v
			}
		}
	}
	return sol
}
