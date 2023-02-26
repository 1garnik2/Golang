package main

import "fmt"

func main() {
	fmt.Println("Please enter an integer from 1 to 10")
	var a int
	fmt.Scan(&a)
	switch a {
	case 1:
		fmt.Println("This is not prime ant not a composite nomber")
	case 2:
		fmt.Println("This is the only even prime number")
	case 3, 5, 7:
		fmt.Println("This is a prime number")
	case 4, 6, 8, 9:
		fmt.Println("This is a composite number")
	default:
		fmt.Println("Number out of range")
	}
}
