package main

import "fmt"

func main() {
	n, err := fmt.Print("Hello")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n)
}
