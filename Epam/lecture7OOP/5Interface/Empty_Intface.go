package main

import "fmt"

func main() {
	var i interface{}
	i = 45
	fmt.Println(i)
	i = "string"
	fmt.Println(i)
	i = true
	fmt.Println(i)
}
