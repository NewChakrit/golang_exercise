package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	sort.Ints(xi)
	fmt.Println(xi)
	// [4 7 12 16 18 42 56 99]

	sort.Strings(xs)
	fmt.Println(xs)
	// [Dr. No James M Moneypenny Q]
}
