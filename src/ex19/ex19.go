package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type Author struct {
	Name string
	Email string
}

type Comment struct {
	Author Author
	Id uint64
	Content string
}

type Article struct {
	Comments []Comment
	Author Author
	Recommends []string
	Id uint64
	Title string
	Content string
}

func main() {

	data := make([]Article, 1) // 값을 저장할 구조체 슬라이스 생성

	data[0].Id = 1
	data[0].Title = "Hello, world!"
	data[0].Author.Name = "Maria"
	data[0].Author.Email = "maria@example.com"
	data[0].Content = "Hello~"
	data[0].Recommends = []string{"John", "Andrew"}
	data[0].Comments = make([]Comment, 1)
	data[0].Comments[0].Id = 1
	data[0].Comments[0].Author.Name = "Andrew"
	data[0].Comments[0].Author.Email = "andrew@hello.com"
	data[0].Comments[0].Content = "Hello Maria"

	buf := new(bytes.Buffer)
	bufEncoder := json.NewEncoder(buf)
	bufEncoder.Encode(data)
	fmt.Println(buf)

	r := strings.NewReader("my request")
	t := make([]byte, 5)
	n, _ := r.Read(t)
	fmt.Println(t)
	fmt.Println(n)

}
