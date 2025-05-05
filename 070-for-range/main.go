package main

import "fmt"

func main() {
	// for range loop
	// -- data structures - slice

	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Println("range over a slice", i, v)
	}
	// result
	//range over a slice 0 42
	//range over a slice 1 43
	//range over a slice 2 44
	//range over a slice 3 45
	//range over a slice 4 46
	//range over a slice 5 47

	// -- data structures - map
	m := map[string]int{
		"James":       42,
		"Monneypenny": 32,
	}
	for k, y := range m {
		fmt.Println("ranging over a map", k, y)
	}
	// result
	// ranging over a map James 42
	//ranging over a map Monneypenny 32
}
