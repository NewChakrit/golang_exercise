package main

import "fmt"

func main() {

	//x := []int{1, 2, 3, 4, 5, 6}
	//defer fmt.Println(foo(x...))
	//
	//y := []int{1, 2, 3, 4, 5}
	//fmt.Println(bar(y))

	defer fmt.Println("3")
	defer fmt.Println("2")
	defer fmt.Println("1")
	fmt.Println(0)

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
