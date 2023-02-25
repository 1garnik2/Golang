package main

import "fmt"

func main() {
	//
	slice := []string{"Pervoe", "Vtoroe", "Tretie"}
	//
	rightcopy := append(make([]string, 0, len(slice)), slice...)

	fmt.Print(rightcopy)

}
