package main

import "fmt"

func main() {
	c := make(chan int)
	// send
	go foo(c)

	// receive
	bar(c)

	fmt.Println("about to exist.")
}

// send
func foo(c chan<- int) {
	c <- 42
}

// receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}

// result
//42
//about to exist.
