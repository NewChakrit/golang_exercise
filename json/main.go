package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll", "Order": "Dasyuromorp"}
]`)

type Animal struct {
	Name  string
	Order string
}

func main() {
	// แปลงเป็น json
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	r, err := json.Marshal(group)
	if err != nil {
		fmt.Println("Error:", err)
	}

	os.Stdout.Write(r)
	// {"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}%

	// แปลงเป็น struct
	var animal []Animal
	err = json.Unmarshal(jsonBlob, &animal)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("\n%+v", animal) // %+v = print key ด้วย
	// [{Name:Platypus Order:Monotremata} {Name:Quoll Order:Dasyuromorp}]%
}
