package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Println(sum2("Todd", 1, 2, 3, 4, 5))
}

func sum(li ...int) int {
	fmt.Println(li)
	fmt.Printf("%T\n", li)

	sum := 0
	for i := 0; i < len(li); i++ {
		sum += li[i]
	}

	return sum
}

func sum2(name string, li ...int) int {
	sum := 0
	for _, i := range li {
		sum += i
	}

	return sum
}
