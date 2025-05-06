package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 1; i <= 42; i++ {
		x := rand.Intn(6)
		switch x {
		case 1:
			fmt.Printf("interation %v\t x is %v.\n", i, x)
		case 2:
			fmt.Printf("interation %v\t x is %v.\n", i, x)
		case 3:
			fmt.Printf("interation %v\t x is %v.\n", i, x)
		case 4:
			fmt.Printf("interation %v\t x is %v.\n", i, x)
		case 5:
			fmt.Printf("interation %v\t x is %v.\n", i, x)
		default:
			fmt.Printf("interation %v\t x is %v.\n", i, x)
		}
	}

}
