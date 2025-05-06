package main

import "fmt"

func main() {
	x := 0

	for {
		if x >= 10 {
			break
		}
		fmt.Println("x is ", x)
		x++
	}

	y := 8

	for {
		if y == 0 {
			break
		}
		fmt.Println("y is ", y)
		y--
	}
}
