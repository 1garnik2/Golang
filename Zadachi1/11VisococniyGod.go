package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	if (n%4 == 0 && n%100 != 0) || n%400 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
		n = n - n%4 + 4
	}
	if (n%4 == 0 && n%100 != 0) || n%400 == 0 {
		fmt.Println(n)
	} else {
		fmt.Println(n + 4)
	}

}
