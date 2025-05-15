package main

import (
	"fmt"
	"math"
)

func main() {

	a := ppwinator(2)
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

}

func ppwinator(a float64) func() float64 {
	var x float64
	return func() float64 {
		x++
		return math.Pow(a, x)
	}
}
