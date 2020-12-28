package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"hello.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	n := 0
	s := "Hello Jong"
	n, err = file.Write([]byte(s))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 저장 완료")

	fi, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	var data = make([]byte, fi.Size())
	file.Seek(0, os.SEEK_SET)

	n, err = file.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 읽기 완료")
	fmt.Println(string(data))
}