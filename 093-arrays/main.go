package main

import "fmt"

func main() {

	a := [3]int{42, 43, 44}
	fmt.Println(a) // [42, 43, 44]

	b := [...]string{"Hello", "New!"}
	fmt.Println(b) // [Hello New!]

	var c [2]int
	c[0] = 7
	c[1] = 8
	fmt.Println(c) // [7, 8]

	fmt.Printf("%T\n", a) // [3]int
	fmt.Printf("%T\n", b) // [2]string
	fmt.Printf("%T\n", c) // [2]int

	{
		var ni [7]int
		ni[0] = 42
		fmt.Printf("%#v \t\t\t\t %T\n", ni, ni)
		// [7]int{42, 0, 0, 0, 0, 0, 0}  [7]int

		// array literal
		ni2 := [4]int{55, 56, 57, 58}
		fmt.Printf("%#v \t\t\t\t\t %T\n", ni2, ni2)
		// [4]int{55, 56, 57, 58}   [4]int

		// array literal
		// have compiler count elements
		ns := [...]string{"Chocolate", "Vanilla", "Strawberry"}
		fmt.Printf("%#v \t %T\n", ns, ns)
		// [3]string{"Chocolate", "Vanilla", "Strawberry"}  [3]string
		fmt.Println(len(ni))  // 7
		fmt.Println(len(ni2)) // 4
		fmt.Println(len(ns))  // 3

	}
}
