package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.NumCPU())
}

func lowercase() {
	for count := 0; count < 3; count ++ {
		for r := 'a'; r <= 'z'; r++ {
			fmt.Printf("%c ", r)
		}
		fmt.Println()
	}
}

func uppercase() {
	for count := 0; count < 3; count ++ {
		for r := 'A'; r <= 'Z'; r++ {
			fmt.Printf("%c ", r)
		}
		fmt.Println()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	go func() {
		lowercase()
		wg.Done()
	}()

	go func() {
		uppercase()
		wg.Done()
	}()

	fmt.Println("Waiting To Finish")

	wg.Wait()

	fmt.Println("\nTerminating Program")
}