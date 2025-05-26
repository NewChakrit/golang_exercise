package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Eem",
		Last:  "Red",
		Age:   29,
	}
	p2 := person{
		First: "James",
		Last:  "Bond",
		Age:   68,
	}

	people := []person{p1, p2}

	v, err := json.Marshal(people)
	if err != nil {
		fmt.Println("Err: ", err)
	}
	fmt.Println(string(v))
	// [{"First":"Eem","Last":"Red","Age":29},{"First":"James","Last":"Bond","Age":68}]
	os.Stdout.Write(v)
	// [{"First":"Eem","Last":"Red","Age":29},{"First":"James","Last":"Bond","Age":68}]%
}
