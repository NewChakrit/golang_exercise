package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Printf("My name's %s and I'm %v.\n", p.first, p.age)
}

func main() {
	p1 := person{
		first: "New",
		age:   28,
	}

	p1.speak() // My name's New and I'm 28.

}
