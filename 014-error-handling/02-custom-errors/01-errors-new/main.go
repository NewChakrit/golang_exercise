package main

import (
	"errors"
	"log"
)

func main() {
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate math: sqare root if negative number.")
	}

	//  2025/06/08 15:36:44 norgate math: sqare root if negative number.
	//	exit status 1

	return 42, nil
}
