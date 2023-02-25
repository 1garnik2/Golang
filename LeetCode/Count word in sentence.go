// Считает максимальное кол-во слов в предложении
package main

import "fmt"

func main() {
	var slice = []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	fmt.Println(mostWordsFound(slice))
}

func mostWordsFound(sentences []string) int {
	var lens, count int
	for k := 0; k < len(sentences); k++ {
		lens = 0
		for _, j := range sentences[k] {
			if j == 32 {
				lens += 1

			}
			if lens > count {
				count = lens
			}
		}
	}

	return count + 1
}
