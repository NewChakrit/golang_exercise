package main

import (
	"fmt"
	"github.com/GoesToEleven/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()

	// also like this
	fmt.Println(s1)
	fmt.Println(s2)
}
