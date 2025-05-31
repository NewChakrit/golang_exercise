package main

import "fmt"

func doSomeThing(x int) int {
	// Does something, already built
	return x * 5
}

func main() {
	ch := make(chan int)
	go func() {
		ch <- doSomeThing(2)
	}()
	fmt.Println(<-ch) // 10
}
