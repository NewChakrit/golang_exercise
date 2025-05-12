package main

import "fmt"

func main() {

	i, s := bar()

	fmt.Println(foo())
	fmt.Println(i)
	fmt.Println(s)
}

func foo() int {
	return 5
}

func bar() (int, string) {
	return 3, "Todd"
}
