package main

import (
	"fmt"
	"sync/atomic"
)

func main() {

	var j int64
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&j, 20)
	}
	fmt.Println(j)
}
