package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []float64{3, 1, 4, 2}
	fmt.Println(madianOne(n))
	fmt.Println("after madianOne", n)

	x := []float64{3, 1, 4, 2}
	fmt.Println(madianTwo(x))
	fmt.Println("after madianTwo", x)
}

func madianOne(x []float64) float64 {
	sort.Float64s(x)

	i := len(x) / 2
	if len(x)%2 == 1 {
		return x[i/2]
	}
	return (x[i-1] + x[i]) / 2
}

func madianTwo(x []float64) float64 {
	// allocate a new underlying array
	n := make([]float64, len(x))
	copy(n, x)

	sort.Float64s(n)

	i := len(n) / 2
	if len(n)%2 == 1 {
		return x[i/2]
	}
	return (n[i-1] + n[i]) / 2
}
