package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		go func(index int) {
			fmt.Printf("I am the %dth gorutine\n", index)

		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}
