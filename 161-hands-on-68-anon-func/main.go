package main

import "fmt"

func main() {
	func(name string) {
		fmt.Printf("Hello! My name's %v.", name)
	}("New")
}
