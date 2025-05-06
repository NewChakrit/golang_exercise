package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 10; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("interation %v x = %v\n", i, x)
		}
	}

}
