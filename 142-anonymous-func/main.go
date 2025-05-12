package main

import "fmt"

func main() {
	func() {
		fmt.Println("anonymous func ran")
	}()

	func(s string) {
		fmt.Println("This is my name", s)
	}("Todd") // This is my name Todd
}

// a named function with an identifier
// func (r receive) identifier(p parameter(s)) (r return(s)) { code }

// an anonymous function
// func  (p parameter(s)) (r return(s)) { code }
