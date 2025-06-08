package main

import (
	"fmt"
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
		return 0, fmt.Errorf("norgate math again: sqare root of negative number: %v", f)
	}
	// fmt.Errorf = return errors.New(Sprintf(format, a...))

	//  resutl:
	//	2025/06/08 15:45:07 norgate math again: sqare root of negative number: -10
	//  exit status 1

	return 42, nil
}
