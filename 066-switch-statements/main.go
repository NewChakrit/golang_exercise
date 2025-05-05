package main

import (
	"fmt"
)

func main() {
	//x := rand.Intn(50)

	x := 40

	switch {
	case x < 42:
		fmt.Println("x is LESS THAN 42")
	case x == 42:
		fmt.Println("x is EQUAL 42")
	case x > 42:
		fmt.Println("x is GREAT THAN 42")
	default:
		fmt.Println("This is the default case for x")
	}

	switch x {
	case 40:
		fmt.Println("x is 40")
	case 41:
		fmt.Println("x is 41")
	case 42:
		fmt.Println("x is 42")
	case 43:
		fmt.Println("x is 43")
	default:
		fmt.Println("This is the default case for x")
	}

	// no default fallthrough

	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough // เพื่อบังคับให้โปรแกรม ดำเนินการเข้า case ถัดไป
	case 41:
		fmt.Println("printed because of fallthrough: x is 41")
		fallthrough
	case 42:
		fmt.Println("printed because of fallthrough: x is 42")
		fallthrough
	case 43:
		fmt.Println("printed because of fallthrough: x is 43")
		fallthrough
	default:
		fmt.Println("printed because of fallthrough: This is the default case for x")
	}
}
