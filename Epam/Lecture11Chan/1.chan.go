package main

import "fmt"

func main() {
	chA := make(chan int)
	chB := make(chan string)

	go func() { chA <- 10 }()
	go func() { chB <- "Hello" }()
	if <-chA == 10 {
		fmt.Println(<-chB)
	}

}
