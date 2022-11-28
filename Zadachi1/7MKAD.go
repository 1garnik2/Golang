package main

import "fmt"

func main() {
	var v, t int32
	fmt.Scan(&v, &t)
	fmt.Println((v*t%109 + 109) % 109)
}
