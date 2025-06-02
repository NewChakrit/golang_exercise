package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go pupolate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func pupolate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	const goroutines = 10
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			for v := range c1 {
				func(v2 int) {
					c2 <- timeConsumingWotk(v2)
				}(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWotk(n int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
