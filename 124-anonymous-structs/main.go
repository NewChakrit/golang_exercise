package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
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
	//struct { first string; last string; age int }
	//main.person

}
