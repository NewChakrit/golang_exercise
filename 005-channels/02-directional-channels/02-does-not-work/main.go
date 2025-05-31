package main

func main() {
	c := make(chan<- int, 2)

	go func() {
		c <- 42
		c <- 43
	}()

	//	fmt.Println(<-c) // invalid operation: cannot receive from send-only channel c (variable of type chan<- int)
	//	fmt.Println(<-c)
}
