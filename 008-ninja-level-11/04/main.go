package main

import (
	"fmt"
	"log"
)

type sqrError struct {
	lat  string
	long string
	err  error
}

func (se sqrError) Error() string {
	return fmt.Sprintf("mat error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.11)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		//e := errors.New("more coffee needed")
		e := fmt.Errorf("more coffee needed - value was %v", f)
		return 0, sqrError{"50.2289", "99.4656 W", e}
	}

	return 42, nil
}
