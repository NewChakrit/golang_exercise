package main

import "fmt"

func main() {
	x := func() {
		fmt.Println("anonymous func ran")
	}

	y := func(s string) {
		fmt.Println("This is my name", s)
	}

	x()
	y("Todd")
}
