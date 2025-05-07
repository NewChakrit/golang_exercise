package main

import (
	"fmt"
)

func main() {
	xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("xs - %v\n", xs)
	fmt.Println("--------------")

	// delete 4
	xs = append(xs[:3], xs[4:]...)
	fmt.Printf("xs - %v\n", xs)

}
