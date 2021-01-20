package main

import (
	//"bufio"
	"fmt"
	//"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("hello.txt")
	check(err)
	fmt.Println(string(dat))

	f, err := os.Open("hello.txt")
	check(err)
	b1 := make([]byte, 15)
	_, err = f.Read(b1)
	check(err)
	fmt.Println(string(b1))

}