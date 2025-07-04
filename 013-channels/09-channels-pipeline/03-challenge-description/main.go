package main

import "fmt"

func main() {
	c := factorail(4)
	for n := range c {
		fmt.Println(n)
	}

}

func factorail(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()

	return out
}

/*
	CHALLENGE #1:
-- Change the code above to execute 100 factorial computations concurrently and in parallel.
-- Use the "pipeline" pattern to accomplish this

POST To Discussion:
-- What realization did you have about working with concurrent code when building your solution?
-- eg, what insights did you have which helped you understand working with concurrency?
 */