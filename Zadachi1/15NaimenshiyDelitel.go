package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	f := true
	for x := 2; x*x <= n; x++ {
		if n%x == 0 {
			fmt.Print(x)
			f = false
			break
		}
	}
	if f {
		fmt.Print(n)
	}

}
