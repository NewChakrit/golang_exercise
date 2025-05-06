package main

import "fmt"

func main() {
	x := 0

	for x < 10 {
		fmt.Println("loop1: x is ", x)
		x++
	}

	y := 10
	for y >= 0 {
		fmt.Println("loop2: y is ", y)
		y--
	}
}
