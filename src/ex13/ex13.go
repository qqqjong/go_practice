package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	step = iota
	any
	gold
)

func main() {
	n, err := strconv.Atoi(os.Args[1])

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered %v\n", r)
		}
	}()

	if err != nil {
		panic(err)
	}

	cnt := make(map[int]int)

	for y:=2020; y<2020+n; y++ {
		fmt.Printf("%d-12-25: ", y)
		switch func(y int) int {
			acc := time.Friday
			for i := 2021; i<=y; i++ {
				if i%400 == 0 || i%4 == 0 && i%100 != 0 {
					acc += 2
				} else {
					acc ++
				}
			}
			switch acc % 7 {
			case time.Monday, time.Friday:
				return gold
			case time.Tuesday, time.Thursday:
				return step
			default:
				return any
			}
		}(y) {
		case step:
			cnt[step]++
			fmt.Println("징검다리")
		case gold:
			cnt[gold]++
			fmt.Println("황금")
		default:
			fmt.Println()
		}
	}
	fmt.Printf("합계 황금: %d, 징검다리 %d\n", cnt[gold], cnt[step])
}