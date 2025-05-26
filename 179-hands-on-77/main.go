package main

import "fmt"

type person struct {
	first string
}

func fly(s *string) {

}

func main() {
	p := person{"Max"}
	fmt.Println(p)

	p = changeName(p, "Jim")
	fmt.Println(p)

	changeNamePointer(&p, "Sam")
	fmt.Println(p)
}

func changeName(p person, s string) person {
	p.first = s
	return p
}

func changeNamePointer(p *person, s string) {
	p.first = s
}
