package main

import "fmt"

func main() {
	c := incrementor()

	for n := range puller(c) {
		fmt.Println(n)
	}
}

func incrementor() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}

func puller(c <-chan int) <-chan int {
	out2 := make(chan int)

	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out2 <- sum
		close(out2)
	}()

	return out2
}
