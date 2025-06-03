package main

import "fmt"

func main() {
	fmt.Println(mySum(2, 3))
	fmt.Println(mySum(4, 7))
	fmt.Println(mySum(12, 11, 10))
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
