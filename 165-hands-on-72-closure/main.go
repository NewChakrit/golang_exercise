package main

import (
	"fmt"
	"math"
)

func main() {

	a := powinator(2)
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

}

func powinator(a float64) func() float64 {
	var x float64
	return func() float64 {
		x++
		return math.Pow(a, x)
	}
}
