package main

import "fmt"

func main() {
	am := map[string]int{
		"Todd":   42,
		"Henry":  16,
		"Padget": 14,
	}

	fmt.Println(am)
	// map[Henry:16 Padget:14 Todd:42]

	fmt.Printf("%#v\n", am)
	// map[string]int{"Henry":16, "Padget":14, "Todd":42}

	for k, v := range am {
		fmt.Printf("key: %v  value: %v\n", k, v)
	}
	// key: Henry  value: 16
	// key: Padget  value: 14
	// key: Todd  value: 42

	fmt.Printf("The age of Henry was %v.\n", am["Henry"])
	// The age of Henry was 16.

	fmt.Println("------------")

	bm := make(map[string]int)
	bm["Lisa"] = 22
	bm["Bob"] = 35

	fmt.Println(bm)
	// map[Bob:35 Lisa:22]

	fmt.Printf("%#v\n", bm)
	// map[string]int{"Bob":35, "Lisa":22}

	fmt.Println(len(bm)) // 2
}
