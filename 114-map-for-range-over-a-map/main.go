package main

import "fmt"

func main() {
	am := map[string]int{
		"Todd":   42,
		"Henry":  16,
		"Padget": 14,
	}

	fmt.Println(am)
	fmt.Printf("%#v\n", am)

	fmt.Println("------------")

	bm := make(map[string]int)
	bm["Lisa"] = 22
	bm["Bob"] = 35
	bm["George"] = 78

	fmt.Println(bm)
	fmt.Printf("%#v\n", bm)
	fmt.Println(len(bm)) // 2

	for k, v := range bm {
		fmt.Printf("%v is %v years old.\n", k, v)
	}

	for _, v := range bm {
		fmt.Println(v)
	}

	for k := range bm {
		fmt.Println(k)
	}

	for i := range bm {
		fmt.Println(i) // Key
	}

}
