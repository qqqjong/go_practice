package main

import (
	"io/ioutil"
	"time"
	"log"
	"os"
	"fmt"
)

func CheckOne() {
	start := time.Now()
	data, err := ioutil.ReadFile("myfile")
	if err != nil {
		panic(err)
	}
	fixedData := string(data)[:20]
	fmt.Println(fixedData)
	elapsed := time.Since(start)
	log.Printf("Function One took %s", elapsed)
}

func CheckTwo() {
	start := time.Now()
	file, err := os.Open("myfile")
	if err != nil {
		panic(err)
	}
	data := make([]byte, 20)
	file.Read(data)
	fmt.Println(string(data))
	elapsed := time.Since(start)
	log.Printf("Function Two took %s", elapsed)
}

func main() {
	CheckOne()
	CheckTwo()
}