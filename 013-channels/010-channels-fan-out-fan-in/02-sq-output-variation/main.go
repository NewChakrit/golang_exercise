package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that read form in.
	c1 := sq(in)
	c2 := sq(in)

	// Consume the merged output from c1 and c2.
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)

	fmt.Printf("TYPE: %T", cs)

	wg.Add(len(cs))
	for _, c := range cs {
		go func(ch <-chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done. This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
