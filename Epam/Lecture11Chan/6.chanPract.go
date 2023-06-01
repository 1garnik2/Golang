package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg2 := sync.WaitGroup{}

	vol := 10
	chDoWork := make(chan int, vol/2)
	chDoAnotherWork := make(chan int, vol/2)
	done := make(chan bool)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go DoWork(chDoWork, i, &wg)
		wg2.Add(1)
		go DoWork(chDoAnotherWork, i, &wg2)
	}

	go checkIfDone(&wg, &wg2, done)

	for {
		select {
		case i := <-chDoWork:
			fmt.Printf("Doing some work %d\n", i)
		case i := <-chDoAnotherWork:
			fmt.Printf("Doing some another work %d\n", i)
		case <-done:
			fmt.Println("Done")
			return
		default:
			fmt.Println("Idling")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func DoWork(ch chan<- int, i int, wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)
	ch <- i
	wg.Done()
}

func checkIfDone(wg *sync.WaitGroup, wg2 *sync.WaitGroup, done chan<- bool) {
	wg.Wait()
	wg2.Wait()

	done <- true
}
