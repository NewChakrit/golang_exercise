package main

import "fmt"

func main() {
	// Set up the pipeline and consume the output.
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n) // 16 then 81
	}

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
		close(out2)
	}()

	return out2
}
