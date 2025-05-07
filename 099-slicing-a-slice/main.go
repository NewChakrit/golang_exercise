package main

import (
	"fmt"
)

func main() {
	xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("xs - %v\n", xs)
	fmt.Println("--------------")

	fmt.Printf("xi - %v\n", xs[:3])
	fmt.Println("--------------")

	fmt.Printf("xo - %v\n", xs[4:])
	fmt.Println("--------------")

	fmt.Printf("xo - %v\n", xs[:])
	fmt.Println("--------------")

}
