// вывод массива разными способами
package main

import "fmt"

func main() {
	mas1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(mas1)

	mas2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("%d", mas2)

	fmt.Println()
	mas3 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", mas3[i])
	}

	fmt.Println()
	mas4 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(mas4); i++ {
		fmt.Printf("%d ", mas4[i])
	}
	fmt.Println()
	mas5 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range mas5 {
		fmt.Print(mas5[i], " ")
	}
	fmt.Println()
	mas6 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, item := range mas6 {
		fmt.Print(item, " ")
	}
}
