package main

import "fmt"

func main() {
	// Ex

	c := make(chan int)
	// send
	go func() {
		c <- 42
		close(c)
	}()

	// receive
	v, ok := <-c

	fmt.Println(v, ok) // 42 true
	v, ok = <-c
	fmt.Println(v, ok) // 0 false

	// ---------------------- //

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("about to exit.")

	// result:
	//the value received from the even channel: 0
	//the value received from the odd channel: 1
	//the value received from the even channel: 2
	//the value received from the odd channel: 3
	//the value received from the even channel: 4
	//the value received from the odd channel: 5
	//the value received from the even channel: 6
	//the value received from the odd channel: 7
	//the value received from the even channel: 8
	//the value received from the odd channel: 9
	//from comma ok 0
	//about to exit.
}

// send channel
func send(even, odd, quit chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

// receive channel
func receive(even, odd, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("the value received from the even channel:", v)
		case v := <-odd:
			fmt.Println("the value received from the odd channel:", v)
		case i, ok := <-quit: // เพื่อตรวจสอบว่าถูกปิดหรือไม่
			if !ok {
				fmt.Println("from comma ok", i)
				return
			} else {
				fmt.Println("from comma ok", i)
			}

		}
	}
}
