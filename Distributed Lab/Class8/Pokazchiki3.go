package main

import "fmt"

func Vtrichy(x *int) {
	*x = (*x) + (*x) + (*x)
}

func main() {
	a := 2
	fmt.Printf("input:%d\n", a)
	Vtrichy(&a)
	fmt.Printf("output:%d", a)
}
