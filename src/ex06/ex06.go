package main

import (
	"fmt"
	"runtime"
)

type count struct {
	num int
}

func (c *count) increment() {
	c.num += 1
}

func (c *count) result() {
	fmt.Println(c.num)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := count{num: 0}
	done := make(chan bool)

	for i := 1; i <= 10000; i++ {
		go func() {
			c.increment()
			done <- true
			runtime.Gosched()
		}()
	}

	for i := 1; i <= 10000; i++ {
		<-done
	}

	c.result()
}
