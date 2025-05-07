package main

import "fmt"

func main() {
	xs := []int{42, 43, 44}
	fmt.Println(xs)
	fmt.Println("--------------")

	// variadic paramater
	xs = append(xs, 45, 46, 47, 48)
	fmt.Println(xs)
	fmt.Println("--------------")
}
