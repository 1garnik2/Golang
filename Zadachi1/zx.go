package main

import "fmt"

func main() {
	var i int32
	fmt.Scan(&i)
	a := i / 1000
	b := (i / 10) % 10
	c := i % 10
	fmt.Println(a, b, c)
	//if a == b || b == c || a == c {
	//	fmt.Println("NO")
	//} else {
	//	fmt.Println("YES")
	//}
}
