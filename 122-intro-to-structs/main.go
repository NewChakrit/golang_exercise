package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
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

	fmt.Printf("%T \t %#v\n", p1, p1)

	fmt.Println(p1.first, p1.last, p1.age)
}
