package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrNorgateMath = errors.New("norgate math: sqare root if negative number.")

func main() {
	fmt.Printf("%T\n", ErrNorgateMath) // *errors.errorString

	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorgateMath
	}

	return 42, nil
}
