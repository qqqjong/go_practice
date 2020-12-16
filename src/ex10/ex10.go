package main

import "fmt"

type Person struct {
	name string
	age int
}

type Student struct {
	Person
	school string
	grade int
}

func (p *Person) greeting() {
	fmt.Println("Hi~")
}

func (s *Student) greeting() {
	fmt.Println("Hi~ Student~")
}

func main() {
	var s Student
	s.Person.greeting()
	s.greeting()
}

