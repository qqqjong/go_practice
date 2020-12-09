package main

import "fmt"

type Car struct {
	val string
}

func MakeTire(carChan chan Car, outChan chan Car) {
	car := <-carChan
	car.val += "Tire, "

	outChan <- car
}

func MakeEngine(carChan chan Car, outChan chan Car) {
	car := <-carChan
	car.val += "Engine, "

	outChan <- car
}

func main() {
	chan1 := make(chan Car)
	chan2 := make(chan Car)
	chan3 := make(chan Car)

	go MakeTire(chan1, chan2)
	go MakeEngine(chan2, chan3)

	chan1 <- Car{val: "Car1: "}

	result := <-chan3
	fmt.Println(result)
}
