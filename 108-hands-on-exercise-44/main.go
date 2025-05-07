package main

import "fmt"

func main() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	exOne := a[:5]
	fmt.Printf("b: %T \t %v\n", exOne, exOne)
	exTwo := a[5:]
	fmt.Printf("b: %T \t %v\n", exTwo, exTwo)
	exThree := a[2:7]
	fmt.Printf("b: %T \t %v\n", exThree, exThree)
	exFour := a[1:6]
	fmt.Printf("b: %T \t %v\n", exFour, exFour)
}
