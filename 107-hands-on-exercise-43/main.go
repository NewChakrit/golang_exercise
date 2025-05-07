package main

import "fmt"

func main() {
	a := []int{}
	a = append(a, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51)

	for _, v := range a {
		fmt.Printf("value: %v \t type: %T\n", v, v)
	}

	fmt.Printf("%T \t %#v\n", a, a)
	fmt.Printf("%T \t %v", a, a)

}
