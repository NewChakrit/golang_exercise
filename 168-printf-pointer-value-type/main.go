package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)  //42
	fmt.Println(&x) //0x1400000e0e0

	fmt.Printf("%v\t%T\n", &x, &x) //0x14000104020   *int
}
