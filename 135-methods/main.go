package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("I'm ", p.first)
}

func main() {
	p1 := person{
		first: "New",
	}

	p2 := person{
		first: "Jenny",
	}

	p1.speak()
	p2.speak()
}
