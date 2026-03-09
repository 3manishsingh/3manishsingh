package main

import "fmt"

// Define an interface
type Speaker interface {
	Speak() string
}

// Define a struct
type Person struct {
	Name string
}

// Implement the interface for Person
type Dog struct {
	Breed string
}

func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

func (d Dog) Speak() string {
	return "Woof!"
}

func main() {
	var s Speaker

	p := Person{Name: "Manish"}
	d := Dog{Breed: "Labrador"}

	s = p
	fmt.Println(s.Speak())

	s = d
	fmt.Println(s.Speak())
}
