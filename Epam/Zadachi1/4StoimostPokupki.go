package main

import "fmt"

func main() {
	var a, b, n int32
	fmt.Scanf("%d %d %d", &a, &b, &n)
	s := (a*100 + b) * n
	fmt.Println(s/100, s%100)

}
