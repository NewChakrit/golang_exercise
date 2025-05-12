package main

import (
	"fmt"
	"math"
)

type Square struct {
	length int
	width  int
}

type Circle struct {
	radius float64
}

func (c Circle) Area() int {
	v := int(math.Pi * math.Pow(c.radius, 2))
	return v
}

func (s Square) Area() int {
	v := s.length * s.width
	return v
}

type Shape interface {
	Area() int
}

func info(s Shape) int {
	return s.Area()
}

func main() {
	s1 := Square{
		length: 4,
		width:  3,
	}

	c1 := Circle{
		radius: 5,
	}

	fmt.Println(info(s1))
	fmt.Println(info(c1))
}
