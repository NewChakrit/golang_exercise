package main

import "fmt"

func main() {
	ten := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}

	xi := make([]string, 0, 10)

	for i, _ := range ten {
		xi = append(xi, ten[i])
	}

	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))

	xs := make([]string, 0, 10)
	fmt.Println(xs)
	fmt.Println(len(xs))
	fmt.Println(cap(xs))
}
