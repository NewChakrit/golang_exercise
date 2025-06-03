package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("norgate math again: square root of negative number: %v", f)
	}
	return 42, nil
}

//result
//2025/06/03 15:30:10 norgate math again: square root of negative number: -10
//exit status 1
