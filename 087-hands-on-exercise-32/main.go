package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	age := m["James"]
	fmt.Println("That age of Bond", age)

	age = m["Q"]
	fmt.Println("That age of", age)

	if v, ok := m["Q"]; !ok {
		fmt.Println("There is no Q", v)
	}
}
