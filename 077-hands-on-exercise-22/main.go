package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(11)
	y := rand.Intn(11)

	fmt.Printf("The value of x is %v\n", x)
	fmt.Printf("The value of y is %v\n", y)

	if x < 4 && y < 4 {
		fmt.Println("x and y are both less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("x and y are both greater than 4")
	} else if x >= 4 && y <= 6 {
		fmt.Println("x is greater than or equal to 4 and less than or equal 6")
	} else if y != 5 {
		fmt.Println("y is not 5")
	} else {
		fmt.Println("none of the previous cases were met")
	}

}
