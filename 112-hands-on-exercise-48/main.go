package main

import "fmt"

func main() {
	a := []string{"James", "Bond", "Shaken, not stirred"}
	b := []string{"Miss", "Monneypenny", "I'm 008."}

	c := make([][]string, 0)
	c = append(c, a, b)
	fmt.Println(c)

	// or
	d := [][]string{a, b}
	fmt.Println(d)

	for i, v := range c {
		fmt.Println(i, v)
		for a, b := range v {
			fmt.Println(a, b)
		}
	}
}
