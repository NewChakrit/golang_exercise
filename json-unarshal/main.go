package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age"`
}

func main() {
	s := `[{"First":"Eem","Last":"Red","Age":29},{"First":"James","Last":"Bond","Age":68}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)  // string
	fmt.Printf("%T\n", bs) // []uint8

	// people := []person{}
	var people []person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n%+v\n", people)
	// [{Eem Red 29} {James Bond 68}]

	for i, v := range people {
		fmt.Println("Person no.", i)
		fmt.Printf("%v\n", v.First)
		fmt.Printf("%v\n", v.Last)
		fmt.Printf("%v\n", v.Age)
	}
}
