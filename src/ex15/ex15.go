package main

import (
	"encoding/json"
	"fmt"
)

type Response1 struct {
	Page int
	Fruits []string
}

type Response2 struct {
	Page int `json:"page""`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(bolB)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(intB)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(3.14)
	fmt.Println(fltB)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(strB)
	fmt.Println(string(strB))

	slcD := []string{"aaa", "bbb", "ccc"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(slcB)
	fmt.Println(string(slcB))

	mapD := map[string]int{"aaa":1, "bbb":2, "ccc":3}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(mapB)
	fmt.Println(string(mapB))

	res1D := &Response1{
		Page: 1,
		Fruits: []string{"apple", "banana", "kiwi"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(res1B)
	fmt.Println(string(res1B))

	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(res2B)
	fmt.Println(string(res2B))

	// =========================================

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	fmt.Println(dat)

}