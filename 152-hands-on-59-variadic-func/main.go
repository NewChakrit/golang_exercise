package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(foo(x...))

	y := []int{1, 2, 3, 4, 5}
	fmt.Println(bar(y))
}

func foo(ii ...int) int {
	sum := 0
	for _, v := range ii {
		sum += v
	}

	return sum
}

func bar(ii []int) int {
	sum := 0
	for _, v := range ii {
		sum += v
	}

	return sum
}
