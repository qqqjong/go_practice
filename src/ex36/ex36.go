package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct {
	Name string
	IsMammal bool
	PackFactor int
}

func (d *Dog) Speak() {
	fmt.Printf(
		"Woof! My name is %s, it is %t I am a mammal with a pack factor of %v",
		d.Name,
		d.IsMammal,
		d.PackFactor,
	)
}

type Cat struct {
	Name string
	IsMammal bool
	ClimbFactor int
}

func (c *Cat) Speak() {
	fmt.Printf(
		"Meow! My name is %s, it is %t I am a mammal with a climb factor of %v",
		c.Name,
		c.IsMammal,
		c.ClimbFactor,
	)
}

func main() {
	speakers := []Speaker{
		&Dog{
			Name: "Fido",
			IsMammal: true,
			PackFactor: 5,
		},
		&Cat{
			Name: "Meme",
			IsMammal: true,
			ClimbFactor: 10,
		},
	}
}