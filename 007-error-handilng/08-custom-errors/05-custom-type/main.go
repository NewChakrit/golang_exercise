package main

import (
	"fmt"
	"log"
)

type norgateMathError struct {
	lat  string
	long string
	err  error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math again: square root of negative number: %v", f)
		return 0, norgateMathError{"50.2289", "99.4656 W", nme}
	}
	return 42, nil
}

// 2025/06/03 15:35:41 a norgate math error occured: 50.2289 99.4656 W norgate math again: square root of negative number: -10.23
//exit status 1
