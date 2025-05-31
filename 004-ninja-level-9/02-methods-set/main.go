package main

import (
	"fmt"
)

type person struct {
	first string
}

func (p *person) speak() string {
	return fmt.Sprintln("Name:", p.first)
}

type human interface {
	speak() string
}

func saySomthing(h human) {
	fmt.Println(h.speak())
}

func main() {
	p := person{"New"}
	saySomthing(&p)
}
