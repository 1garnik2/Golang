package main

import "fmt"

func main() {
	ar := []int{3, 4, 1, 2, 6, 5, 7, -1, 0}
	for i := 0; i < len(ar); i++ {
		for j := i; j < len(ar); j++ {
			if ar[i] > ar[j] {
				tmp := ar[i]
				ar[i] = ar[j]
				ar[j] = tmp
			}
		}
	}
	fmt.Println(ar)
}
