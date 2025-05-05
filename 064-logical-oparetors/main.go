package main

import "fmt"

func main() {
	x := 40
	y := 5

	if x < 42 && y < 42 {
		fmt.Println("both are less than the meaning of life")
	}

	if x > 30 || x < 42 {
		fmt.Println("x is getting close to the meaning of life")
	}

	if x != 42 || y != 10 {
		fmt.Println("x is not 42 & y is not 10")
	}

	//
}
