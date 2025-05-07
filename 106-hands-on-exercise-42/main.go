package main

import "fmt"

func main() {
	a := [5]int{}
	for i := 0; i < 5; i++ {
		a[i] = i
	}

	for _, v := range a {
		fmt.Printf("type: %T \t value: %v\n", v, v)
	}

}
