package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(251)

	fmt.Printf("x is %v\n", x)

	if x <= 100 {
		fmt.Println("x is between 0 and 100")
	} else if x > 100 && x < 201 {
		fmt.Println("x is between 101 and 200")
	} else {
		fmt.Println("x is between 201 and 250")
	}

}
