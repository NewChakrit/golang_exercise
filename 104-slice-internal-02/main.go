package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	b := make([]int, 6)
	copy(b, a)

	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("-----------")

	a[0] = 7
	b[5] = 10
	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("-----------")
	//a  [7 1 2 3 4 5]
	//b  [0 1 2 3 4 10]
}
