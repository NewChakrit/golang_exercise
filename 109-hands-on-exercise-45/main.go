package main

import "fmt"

func main() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	a = append(a, 52)
	fmt.Printf("a: %v\n", a)

	a = append(a, 53, 54, 55)
	fmt.Printf("a updated: %v\n", a)

	y := []int{56, 57, 58, 59, 60}

	for _, v := range y {
		a = append(a, v)
	}
	fmt.Printf("a + y: %v\n", a)

	z := []int{61, 62, 63, 64}
	a = append(a, z...)
	fmt.Printf("a + z: %v\n", a)

}
