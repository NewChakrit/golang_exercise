package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()

	// FAN OUT
	// Multiple function reading from the same channel until that channel is closed
	// Distribute work across multiple function (ten goroutine) that all read from in.

	c0 := factorial(in)
	c1 := factorial(in)
	c2 := factorial(in)
	c3 := factorial(in)
	c4 := factorial(in)
	c5 := factorial(in)
	c6 := factorial(in)
	c7 := factorial(in)
	c8 := factorial(in)
	c9 := factorial(in)

	//FAN IN
	// multiplex multiple channels onto a single channel
	// merge the channels from c0 through c9 into a single channel

	y := 0
	for n := range merge(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9) {
		y++ // check execute 100 factorial computations
		fmt.Println(y, "\t", n)
	}

}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done. This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()

	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()

	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}

	return total
}

/*
CHALLENGE #1
-- change the code above to execute 100 factorial computations concurrently and in parallel.
-- Use the "fan out / fan in" pattern to accomplish this

CHALLENGE #2
-- While running the factorial computations, try to find how much of your resources are being used.
-- POST the percentage of your resources  being used to this discussion
*/
