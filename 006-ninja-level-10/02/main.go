package main

import "fmt"

func main() {
	cs := make(chan int, 1)

	go func() {
		cs <- 42
	}()

	fmt.Println(<-cs)

	fmt.Printf("--------\n")
	fmt.Printf("cs\t%T\n", cs)
}
