package main

import "fmt"

type person struct {
	first string
}

type secreteAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("I'm ", p.first)
}

func (sa secreteAgent) speak() {
	fmt.Println("I'm a secret agent", sa.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	sa1 := secreteAgent{
		person: person{
			first: "New",
		},
		ltk: true,
	}

	p2 := person{
		first: "Jenny",
	}

	saySomething(sa1)
	saySomething(p2)

}
