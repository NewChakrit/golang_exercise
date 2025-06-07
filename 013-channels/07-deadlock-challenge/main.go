package main

import "fmt"

//func main() {
//	c := make(chan int)
//	go func() {
//		c <- 1
//	}()
//
//	//c <- 1  // result = deadlock
//
//	fmt.Println(<-c)
//}

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

}
