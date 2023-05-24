package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	var j int
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(100 * time.Millisecond)
			j++
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(j)
}
