package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := ioutil.ReadFile("./data.txt")
	check(err)
	fmt.Println(data)
	fmt.Println(string(data))

	f, err := os.Open("./data.txt")
	check(err)

	b1 := make([]byte, 5)
	f.Read(b1)
	fmt.Println(b1)
	fmt.Println(string(b1))

	f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	f.Read(b2)
	check(err)
	fmt.Println(b2)
	fmt.Println(string(b2))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Println(b4)
	fmt.Println(string(b4))

	f.Close()

}
