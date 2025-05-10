package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sum(xi...)) //!!!!!
}

func sum(li ...int) int {
	sum := 0
	for i := 0; i < len(li); i++ {
		sum += li[i]
	}

	return sum
}
