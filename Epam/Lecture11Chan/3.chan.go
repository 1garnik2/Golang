// Односторонний канал
package main

import "fmt"

func main() {
	chBVolume := 10
	chB := make(chan string, chBVolume)

	for i := 0; i < chBVolume; i++ {
		chB <- fmt.Sprintf("Hello chanals g %d", i)
	}

	oneDChanB := giveMeOneDChannel(chB)

	for i := 0; i < chBVolume; i++ {
		fmt.Println(<-oneDChanB)
	}
	// oneDChanB <- "My New String"   // error запись в канал не возможна
}

func giveMeOneDChannel(ch chan string) <-chan string {
	return ch
}
