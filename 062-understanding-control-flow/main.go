package main

import "fmt"

// Run before function main
func init() {
	fmt.Println("Begin Initialization")
}

func main() {
	// SEQUENCE
	fmt.Println("This is the first statment to run")
	fmt.Println("This is the second statment to run")
	x := 40
	y := 5
	fmt.Printf("x=%v \ny=%v \n", x, y)
}
