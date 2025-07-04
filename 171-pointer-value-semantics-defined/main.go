package main

import "fmt"

// value semantics
func addOne(v int) int {
	return v + 1
}

// pointer semantics
func addOneP(v *int) {
	*v++
}

func main() {
	a := 1
	fmt.Println(a)         // 1
	fmt.Println(addOne(a)) // 2
	fmt.Println(a)         // 1

	// pointer semantics
	b := 1
	fmt.Println("value semantics")
	addOneP(&b)
	fmt.Println(b) // 2
}
