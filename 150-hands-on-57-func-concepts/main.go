package main

import "fmt"

func main() {
	h := helloWorld
	h()

	s := sum([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(s)
}

// func (r receiver) identifier(p parameter(s)) (r return(s)) {code}

func helloWorld() {
	fmt.Println("Hello World!")
}

func sum(ii []int) (total int) {
	total = 0
	for _, v := range ii {
		total += v
	}

	return total
}
