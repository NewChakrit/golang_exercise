package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string   `json:"first"`
	Last    string   `json:"last"`
	Age     int      `json:"age"`
	Sayings []string `json:"sayings"`
}

func main() {
	u1 := user{"James", "Bond", 32, []string{"Shaken, not stirred", "Youth is no guarantee of innovation", "In this majesty's royal service"}}
	u2 := user{"Miss", "Moneypenny", 27, []string{"James, it is soo good to see you", "Would you like ne to take care of that for you, James?", "I would really prefer to be a secret agent myself"}}
	u3 := user{"M", "Hmmmm", 54, []string{"Oh, James. you didn't", "Dear god, what has James done not?", "Can someone please tell me where James Bond is?"}}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
