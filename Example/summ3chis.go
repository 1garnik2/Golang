package main

import "fmt"

var a, b, c int

func main() {

	fmt.Scanf("%d %d %d", &a, &b, &c)
	var m = int(a + b + c)
	fmt.Println(m)
}
