package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string   `json:"first"`
	Last    string   `json:"last"`
	Age     int      `json:"age"`
	Sayings []string `json:"sayings"`
}

func main() {
	s := `[{"First":"Eem","Last":"Red","Age":29, "Sayings":["Hi", "HO", "HOO"]},{"First":"James","Last":"Bond","Age":68,"Sayings":["MI", "MO", "MOO"]}]`

	var people []person

	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)

	for _, v := range people {
		fmt.Printf("%v %v %v\n", v.First, v.Last, v.Age)
		for _, vv := range v.Sayings {
			fmt.Printf("\t%v\n", vv)
		}
	}
}
