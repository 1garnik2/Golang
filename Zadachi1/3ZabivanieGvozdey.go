package main

import "fmt"

func main() {
	var l, k int32
	fmt.Scan(&l, &k)
	fmt.Println((l*10 + k - 1) / k)
}
