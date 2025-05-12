package main

import "fmt"

func main() {
	fmt.Println(foo([]int{1, 2, 3, 4}))
}

func foo(ii ...[]int) int {
	sum := 0
	for _, v := range ii {
		for _, v2 := range v {
			sum = sum + v2
		}
	}

	return sum
}
