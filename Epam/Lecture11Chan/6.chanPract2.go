package main

import (
	"fmt"
	"math/rand"
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

	go func(chDoWork chan<- int, chDoAnotherWork chan<- int, wg *sync.WaitGroup, wg2 *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go doWork(chDoWork, i, wg)
			wg2.Add(1)
			go doWork(chDoAnotherWork, i, wg2)
		}
		go checkIfDone(wg, wg2, done)
	}(chDoWork, chDoAnotherWork, &wg, &wg2)
	wg3 := sync.WaitGroup{}
	wg3.Add(1)
	go selectFunc(chDoWork, chDoAnotherWork, done, &wg3)
	fmt.Println("I'am still alive!!!")
	wg3.Wait()

}
func selectFunc(chDoWork <-chan int, chDoAnotherWork <-chan int, done <-chan bool, wg *sync.WaitGroup) {
	for {
		select {
		case i := <-chDoWork:
			fmt.Printf("Doing some work %d...\n", i)
		case i := <-chDoAnotherWork:
			fmt.Printf("Doing some another work %d...\n", i)
		case <-done:
			fmt.Println("Done")
			wg.Done()
			return
		default:
			fmt.Println("Idling...")
			time.Sleep(200 * time.Millisecond)
		}
	}
}
func doWork(ch chan<- int, i int, wg *sync.WaitGroup) {
	time.Sleep(time.Duration(rand.Intn(1000)) * 10 * time.Millisecond)
	ch <- i
	wg.Done()
}

func checkIfDone(wg *sync.WaitGroup, wg2 *sync.WaitGroup, done chan<- bool) {
	wg.Wait()
	wg2.Wait()

	done <- true
}
