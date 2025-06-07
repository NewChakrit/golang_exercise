package main

import "fmt"

func main() {
	// Set up the pipeline.
	c := gen(2, 3)
	out := sq(c)

	// Consume th e output.
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
}

func gen(nums ...int) chan int {
	out1 := make(chan int)
	go func() {
		for _, n := range nums {
			out1 <- n
		}
		close(out1)
	}()

	return out1
}

func sq(in chan int) chan int {
	out2 := make(chan int)
	go func() {
		for n := range in {
			out2 <- n * n
		}
		//close(out2)
	}()

	return out2
}
