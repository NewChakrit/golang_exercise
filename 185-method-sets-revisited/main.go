package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{5}
	//info(c) //  cannot use c (variable of struct type circle) as shape value in argument to info: circle does not implement shape (method area has pointer receiver)
	fmt.Println(c.area()) // 78.53981633974483
}
