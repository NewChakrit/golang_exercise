package main

import "fmt"

func main() {
	x := 83 / 40 // 2
	y := 83 % 40 // 3
	fmt.Println(x, y)

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("even number: ", i)
		} else {
			fmt.Println("odd number: ", i)
		}
	}
}
