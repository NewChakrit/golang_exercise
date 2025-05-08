package main

import "fmt"

func main() {
	b := map[string][]string{}
	b["bond_james"] = []string{"shaken, not strirred", "martinis", "fast cars"}
	b["moneypenny_jenney"] = []string{"intelligence", "literature", "computer science"}
	b["no_dr"] = []string{"cats", "ice cream", "sunsets"}
	b["fleming_ian"] = []string{"straks", "cigars", "espionage"}

	for k, v := range b {
		fmt.Println("----------")
		fmt.Printf("%v\n", k)

		for i, v2 := range v {
			fmt.Printf("%v  %v\n", i, v2)
		}
	}
}
