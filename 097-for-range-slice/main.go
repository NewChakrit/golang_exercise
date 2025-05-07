package main

import "fmt"

func main() {
	xs := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)"}

	for _, v := range xs {
		fmt.Printf("%v\n", v)
	}

	fmt.Println("--------------")
	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])

	fmt.Println("--------------")

	for i := 0; i < len(xs); i++ {
		fmt.Printf("%v\n", xs[i])
	}
}
