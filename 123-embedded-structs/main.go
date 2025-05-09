package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "Jim",
			last:  "Karry",
			age:   52,
		},
		ltk: true,
	}

	p1 := person{
		first: "New",
		last:  "Wen",
		age:   29,
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(sa1)

	fmt.Printf("%T \t %#v\n", p1, p1)

	fmt.Println(p1.first, p1.last, p1.age)

	fmt.Println(sa1.first, sa1.person.first)
}
