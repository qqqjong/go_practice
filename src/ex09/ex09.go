package main

import (
	"fmt"
	"unsafe"
)

type Foo struct {
	w byte
	x byte
	y uint64
}

type Bar struct {
	x byte
	y uint64
	w byte
}

func main() {
	newFoo := new(Foo)
	newBar := new(Bar)

	fmt.Println(unsafe.Sizeof(*newFoo))
	fmt.Println(unsafe.Sizeof(*newBar))
}