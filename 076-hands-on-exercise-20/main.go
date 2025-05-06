package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	fmt.Println("Let's start exercise 076!")
}

func main() {
	x := rand.Intn(3)

	switch x {
	case 1:
		println("x is 1")
	case 2:
		println("x is 2")
	default:
		println("x is 0")
	}

	switch {
	case x == 1:
		println("x is 1")
	case x == 2:
		println("x is 2")
	default:
		println("x is 0")
	}

	ch1 := make(chan int)
	ch2 := make(chan int)

	d1 := time.Duration(rand.Intn(250))
	d2 := time.Duration(rand.Intn(250))

	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 10
	}()

	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 20
	}()

	select {
	case v1 := <-ch1:
		fmt.Println("value from chanel 1", v1)
	case v2 := <-ch2:
		fmt.Println("value from chanel 2", v2)

	}
}
