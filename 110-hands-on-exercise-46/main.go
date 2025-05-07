package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	//y := x[6:] // [48 49 50 51]
	//
	//x = x[:3]
	//x = append(x, y...)

	x = append(x[:3], x[6:]...)
	fmt.Println(x)

	// -------------
}
