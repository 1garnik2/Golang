package main

import (
	"fmt"
	"time"
)

func start() {
	fmt.Println("This in gorutines")
}

func main() {
	go start()
	fmt.Println("Start")
	time.Sleep(100 * time.Millisecond)
	go start()
	fmt.Println("Finish")
}
