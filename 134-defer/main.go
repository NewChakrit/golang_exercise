package main

import "fmt"

func main() {
	defer foo()
	bar()

	//bar
	//foo
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
