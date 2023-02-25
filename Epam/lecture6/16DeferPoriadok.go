package main

import "fmt"

func main() {
	// первый дефер будет последним
	defer finish()
	defer fmt.Println("Program has been started")
	fmt.Println("Program is working")
}
func finish() {
	fmt.Println("Program has been finished")
}
