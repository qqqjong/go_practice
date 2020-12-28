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
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("./data.txt", d1, 0644)
	check(err)

	f, err := os.Create("./data2.txt")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Println(n2)

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered")
	fmt.Println(n4)

	w.Flush()
}