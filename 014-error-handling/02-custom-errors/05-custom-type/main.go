package main

import (
	"fmt"
	"log"
)

type NorgateMathError struct {
	lat, long string
	err       error
}

func (n *NorgateMathError) Error() string {
	return fmt.Sprintf("a noragte math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func Sqrt(f float64) (float64, error) {
	nme := fmt.Errorf("norgate math again: sqare root of negative number: %v", f)

	if f < 0 {
		return 0, &NorgateMathError{
			lat:  "45.6",
			long: "22.78",
			err:  nme,
		}
	}

	//  resutl:
	//2025/06/08 15:54:02 a noragte math error occured: 45.6 22.78 norgate math again: sqare root of negative number
	//exit status 1

	return 42, nil
}
