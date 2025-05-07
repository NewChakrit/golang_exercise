package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	b := a

	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("-----------")

	a[0] = 7
	b[5] = 10
	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("-----------")
	//a  [7 1 2 3 4 10]
	//b  [7 1 2 3 4 10]

	c := b
	c[5] = 21
	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("c ", c)
	fmt.Println("-----------")
	//a  [7 1 2 3 4 21]
	//b  [7 1 2 3 4 21]
	//c  [7 1 2 3 4 21]

}
