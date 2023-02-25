package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	switch a {
	case 1:
		fmt.Println("Do Task 1")
		fallthrough // проваливаеться в нижний кейс
	case 2:
		fmt.Println("Do Task 2")
		fallthrough
	case 3:
		fmt.Println("Do Task 3")

	case 4:
		fmt.Println("Do Task 4")
		fallthrough
	case 5:
		fmt.Println("Do Task 5")
		fallthrough
	case 6:
		fmt.Println("Do Task 6")
		fallthrough
	default:
		fmt.Println("No")
	}
}
