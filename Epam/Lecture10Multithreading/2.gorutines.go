package main

import (
	"fmt"
	"time"
)

func start(number int, sleepTime time.Duration) {
	time.Sleep(sleepTime)
	fmt.Println("In Gorutine", number)
}

func main() {
	for i := 0; i < 10; i++ {
		go start(i, time.Millisecond*500)
	}
	fmt.Println("Started")
	time.Sleep(1 * time.Second)
	fmt.Println("Finished")
}
