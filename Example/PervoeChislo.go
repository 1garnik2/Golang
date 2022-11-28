// первое число из 10000
package main

import "fmt"

func main() {
	var a, b, c, d, e, f int32
	fmt.Scan(&a)
	f = a / 10000
	b = a / 1000
	c = a / 100
	d = a / 10
	e = a % 10
	if f >= 1 {
		fmt.Println(f)
	} else if b >= 1 {
		fmt.Println(b)
	} else if c >= 1 {
		fmt.Println(c)
	} else if d >= 1 {
		fmt.Println(d)
	} else {
		fmt.Println(e)
	}

}
