package main

import "fmt"

type Car struct {
}

type Airplane struct {
}

func (c *Car) ride() {
	fmt.Println("부릉")
}

func (c *Car) sound() {
	fmt.Println("빵빵")
}

func (a *Airplane) ride() {
	fmt.Println("슝~")
}

func (a *Airplane) sound() {
	fmt.Println("후잉")
}

type Traffic interface {
	ride()
	sound()
}

func isTraffic(t Traffic) {
	t.ride()
	t.sound()
}

func main() {
	var car Car
	var airplane Airplane

	isTraffic(&car)
	isTraffic(&airplane)
}
