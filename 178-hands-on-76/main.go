package main

import "fmt"

type dog struct {
	first string
}

func (d dog) run() {
	fmt.Println("My name is", d.first, "and I'm running.")
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walking.")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.walk()
	y.run()
}

func main() {
	d1 := dog{"Henry"}
	youngRun(&d1)

	fmt.Println("===================")
	d2 := dog{"Padget"}
	youngRun(d2)

	fmt.Println("===================")
	d2 = d2.changeName("Minny")
	youngRun(d2)
}

func (d dog) changeName(s string) dog {
	d.first = s

	return d
}
