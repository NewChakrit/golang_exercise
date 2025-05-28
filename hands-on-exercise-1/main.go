package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	First string
	Age   int
}

func main() {
	p1 := User{"New", 28}
	p2 := User{"Moneypenny", 27}
	p3 := User{"M", 24}

	users := []User{p1, p2, p3}

	v, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(string(v))

	os.Stdout.Write(v)

	for _, v := range users {
		fmt.Printf("\n%v: %v", v.First, v.Age)
	}
}
