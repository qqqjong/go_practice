package main

import (
	"fmt"
	"runtime"
	"sync"
)

var ch = make(chan int)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	mutex := sync.Mutex{}

	for i := 1; i <= 1000; i++ {
		go func(num int) {
			ch <- num
		}(i)

	}

	sum := 0
	for i := 1; i <= 1000; i++ {
		go func() {
			mutex.Lock()
			sum += <-ch
			mutex.Unlock()
		}()
	}

	fmt.Println(sum) // 500500
}

