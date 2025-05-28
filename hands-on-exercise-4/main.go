package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 43, 17, 987, 14, 12, 21, 1, 4, 2}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under"}

	sort.Ints(xi)
	fmt.Println(xi)

	sort.Strings(xs)
	fmt.Println(xs)
}
