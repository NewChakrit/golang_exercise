package main

import "fmt"

func main() {
	bm := make(map[string]int)
	bm["Lisa"] = 22
	bm["Bob"] = 35
	bm["George"] = 78

	// Delete
	delete(bm, "Lisa")
	delete(bm, "Tim") // Not thing change

	fmt.Println(bm)        // map[Bob:35 George:78]
	fmt.Println(bm["Tim"]) // 0

	v, ok := bm["Tim"]
	fmt.Printf("%v \t%v\n", v, ok) // 0 false

	u, ok := bm["George"]
	fmt.Printf("%v \t%v\n", u, ok) // 78 true

	if k, ok := bm["Bob"]; ok {
		fmt.Println(k) // 35
	} else {
		fmt.Println("The key doesn't exist.")
	}

	if k, ok := bm["Bob"]; !ok {
		fmt.Println(k)
	} else {
		fmt.Println("The key doesn't exist.") // The key doesn't exist.
	}

}
