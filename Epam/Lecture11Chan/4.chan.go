package main

import "fmt"

func main() {
	ch := make(chan int, 10)

	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) //канал нужно закрывать перед циклом
	for item := range ch {
		fmt.Println(item)
	}

}
