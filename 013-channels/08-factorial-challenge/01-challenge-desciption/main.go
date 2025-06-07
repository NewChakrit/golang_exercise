package main

import "fmt"

func main() {
	f := factorial(4)
	fmt.Println("Total:", f)
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}

	return total
}

/*
CHALLANGE #1
-- USE goroutines and channelds to calculates factorial?

CHALLANGE #2
-- Why might you want to use goroutines and channels to calculate factorial?
---- Formulate your own answer, then post that answer to this discussion area:
---- Read a few to the other answers at the discusstion area to see the reasons of others
*/
