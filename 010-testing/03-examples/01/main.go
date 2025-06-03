package main

import (
	"fmt"
	"github.com/NewChakrit/golang_exercise/010-testing/03-examples/01/acdc"
)

func main() {
	fmt.Println(acdc.Sum(2, 3))
	fmt.Println(acdc.Sum(2, 3, 4, 5, 6, 7, 8, 9))
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
