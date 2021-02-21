package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	mutex := sync.Mutex{}

	var count int

	ch := make(chan int)

	for i := 1; i <= 1000; i++ {
		go func() {
			mutex.Lock()
			count++
			mutex.Unlock()
			ch <- count
			runtime.Gosched()
		}()
	}

	total := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			total += <-ch
		}()
	}
	fmt.Print(total)
}
