package main

import "fmt"

func main() {
	x := doMath(1, 2, add)
	fmt.Println(x) // 3

	y := doMath(42, 16, subtract)
	fmt.Println(y) // 26

	fmt.Printf("%T \n", add)      // func(int, int) int
	fmt.Printf("%T \n", subtract) // func(int, int) int
	fmt.Printf("%T \n", doMath)   // func(int, int, func(int, int) int) int
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}
