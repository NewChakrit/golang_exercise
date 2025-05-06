package main

import "fmt"

func main() {
	for i := 10; i != 0; i-- {
		if i%2 != 0 {
			println("Odd number is ", i)
		}

		fmt.Println(i)
	}
}
