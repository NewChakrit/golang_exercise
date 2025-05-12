package main

import "fmt"

func main() {

	x := foo()
	fmt.Println(x) // 42

	y := bar()
	y()

	fmt.Println(y()) // 43

	fmt.Printf("%T \n", foo) // func() int
	fmt.Printf("%T \n", bar) // func() func() int
	fmt.Printf("%T \n", y)   // func() int
}

func foo() int {
	return 42
}

func bar() func() int {
	return func() int {
		return 43
	}
}
