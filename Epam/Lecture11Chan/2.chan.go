// Буфиризированный канал
package main

import "fmt"

func main() {
	chBVolume := 10
	chB := make(chan string, chBVolume)

	for i := 0; i < chBVolume; i++ {
		chB <- fmt.Sprintf("Hello channels %d", i)
	}
	for i := 0; i < chBVolume; i++ {
		fmt.Println(<-chB)
	}

}
