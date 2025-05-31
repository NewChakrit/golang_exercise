package main

import "fmt"

func main() {
	c := make(<-chan int, 2)

	//go func() {
	//	c <- 42
	//	c <- 43
	//}()

	fmt.Println(<-c)
	fmt.Println(<-c) // invalid operation: cannot send to receive-only channel c (variable of type <-chan int)

}
