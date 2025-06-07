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
