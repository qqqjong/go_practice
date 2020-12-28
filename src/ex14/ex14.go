package main

import (
	s "strings"
	"fmt"
)

func main() {
	fmt.Println("Contains: ", s.Contains("test", "te"))
	fmt.Println("Count: ", s.Count("cookie", "o"))
	fmt.Println("HasPrefix: ", s.HasPrefix("test", "te"))
	fmt.Println("HasSuffix: ", s.HasSuffix("test", "st"))
	fmt.Println("Index: ", s.Index("test", "t"))
	fmt.Println("Join: ", s.Join([]string{"a", "b"}, "-"))
	fmt.Println("Repeat: ", s.Repeat("a", 7))
	fmt.Println("Replace: ", s.Replace("foo", "o", "0", -1))
	fmt.Println("Split: ", s.Split("a-b-c-d-e", "-"))
	fmt.Println("To Lower: ", s.ToLower("TEST"))
	fmt.Println("To Upper: ", s.ToUpper("test"))
	fmt.Println("Len: ", len("Hello"))
	fmt.Println("Char: ", "Hello"[0])
}