package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type foo int

func main() {

	var a foo = 42
	fmt.Printf("%T\n", a)

	// anonymous struct
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "New",
		last:  "Wen",
		age:   29,
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   27,
	}

	//fmt.Println(p1)
	//fmt.Println(p2)

	fmt.Printf("%T\n", p1)
	fmt.Printf("%T\n", p2)
	fmt.Printf("%#v\n", p2)
	//struct { first string; last string; age int }
	//main.person

	p2 = p1
	fmt.Printf("%T\n", p2)
	fmt.Printf("%#v\n", p2)

}
