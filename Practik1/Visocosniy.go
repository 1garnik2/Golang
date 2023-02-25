package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)
	if (i%4 == 0 && i%100 != 0) || (i%400 == 0) {
		fmt.Println("YES")
		for {
			i++
			if (i%4 == 0 && i%100 != 0) || (i%400 == 0) {
				fmt.Println(i)
				break
			}
		}
	} else {
		fmt.Println("NO")
		for {
			i++
			if (i%4 == 0 && i%100 != 0) || (i%400 == 0) {
				fmt.Println(i)

				break
			}

		}
	}

}
