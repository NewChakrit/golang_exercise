package main

import "fmt"

func main() {
	x := 45
	y := 5
	fmt.Printf("x=%v \ny=%v \n", x, y)

	// CONDITION
	// if statements
	// switch statements

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else {
		fmt.Println("equal to, or greater than, the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else if x == 42 {
		fmt.Println("equal than the meaning of life")
	} else {
		fmt.Println("greater than, the meaning of life")
	}

}
