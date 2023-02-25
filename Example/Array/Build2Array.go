package main

import "fmt"

func main() {
	mas := [4][3]int{{1, 2, 3}, {4, 5, 6}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(mas)

	for i := 0; i < 4; i++ { //перебириает рядки
		fmt.Println(mas[i]) //и выводит с нового рядка
	}
	//или с []
	for i := range mas {
		fmt.Println(mas[i])
	}
	//или c []
	for _, i := range mas {
		fmt.Println(i)
	}
	//без[]
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(mas[i][j], " ")
		}
		fmt.Println()
	}

}
