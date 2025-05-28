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
	users := `[{"First": "James", "Last": "Bond", "Age": 32, "Sayings": ["Shaken, not stirred", "You is no guarantee of innovation"]}]`

	var people []person
	err := json.Unmarshal([]byte(users), &people)
	if err != nil {
		fmt.Println("Error:", err)
	}

	//fmt.Println(people)

	for _, v := range people {
		fmt.Printf("First: %v\nLast: %v\nAge: %v\n", v.First, v.Last, v.Age)
		for _, vv := range v.Sayings {
			fmt.Printf("Sayings: %v\n", vv)
		}
	}
}
