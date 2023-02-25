//Считает Х если Х положительные то знач +1
//если отриц. то значение -1

package main

import "fmt"

func main() {
	var arr = [3]string{"--X", "X++", "X++"}
	slice := arr[:]
	fmt.Println(finalValueAfterOperations(slice))
}
func finalValueAfterOperations(operations []string) int {
	count := 0
	for _, v := range operations {
		if v == "--X" || v == "X--" {
			count -= 1
		} else {
			count += 1
		}
	}
	return count
}
