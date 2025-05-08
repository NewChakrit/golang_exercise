package main

import "fmt"

func main() {

	subA1 := []string{"shaken, not strirred", "martinis", "fast cars"}
	subA2 := []string{"intelligence", "literature", "computer science"}
	subA3 := []string{"cats", "ice cream", "sunsets"}

	a := map[string][]string{
		"bond_james":        subA1,
		"moneypenny_jenney": subA2,
		"no_dr":             subA3,
	}

	for k, v := range a {
		fmt.Println("----------")
		fmt.Printf("k: %v  v: %v\n", k, v)

		for i, v2 := range v {
			fmt.Printf("index: %v \t key: %v \t value: %v\n", i, k, v2)
		}
	}

	fmt.Println("===============================")

	b := map[string][]string{}
	b["bond_james"] = []string{"shaken, not strirred", "martinis", "fast cars"}
	b["moneypenny_jenney"] = []string{"intelligence", "literature", "computer science"}
	b["no_dr"] = []string{"cats", "ice cream", "sunsets"}

	for k, v := range b {
		fmt.Println("----------")
		fmt.Printf("%v\n", k)

		for i, v2 := range v {
			fmt.Printf("%v  %v\n", i, v2)
		}
	}
}
