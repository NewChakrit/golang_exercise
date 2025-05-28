package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string   `json:"first"`
	Last    string   `json:"last"`
	Age     int      `json:"age"`
	Sayings []string `json:"sayings"`
}

// First
type byFirst []user

func (a byFirst) Len() int {
	return len(a)
}

func (a byFirst) Less(b, c int) bool {
	return a[b].First < a[c].Last
}

func (a byFirst) Swap(b, c int) {
	a[b], a[c] = a[c], a[b]
}

// Last
type byLast []user

func (a byLast) Len() int {
	return len(a)
}

func (a byLast) Less(b, c int) bool {
	return a[b].Last < a[c].Last
}

func (a byLast) Swap(b, c int) {
	a[b], a[c] = a[c], a[b]
}

// Age
type byAge []user

func (a byAge) Len() int {
	return len(a)
}

func (a byAge) Less(b, c int) bool {
	return a[b].Age < a[c].Age
}

func (a byAge) Swap(b, c int) {
	a[b], a[c] = a[c], a[b]
}

func main() {
	u1 := user{"James", "Bond", 32, []string{"Shaken, not stirred", "Youth is no guarantee of innovation", "In this majesty's royal service"}}
	u2 := user{"Miss", "Moneypenny", 27, []string{"James, it is soo good to see you", "Would you like ne to take care of that for you, James?", "I would really prefer to be a secret agent myself"}}
	u3 := user{"M", "Hmmmm", 54, []string{"Oh, James. you didn't", "Dear god, what has James done not?", "Can someone please tell me where James Bond is?"}}

	users := []user{u1, u2, u3}

	//fmt.Println(users)

	sort.Sort(byFirst(users))
	for _, v := range users {
		sort.Strings(v.Sayings)
		fmt.Printf("%v %v %v \n", v.First, v.Last, v.Age)
		for _, vv := range v.Sayings {
			fmt.Printf("\t%v\n", vv)
		}
	}

	fmt.Println("===============")
	sort.Sort(byLast(users))
	for _, v := range users {
		fmt.Printf("%v %v %v \n", v.First, v.Last, v.Age)
		for _, vv := range v.Sayings {
			fmt.Printf("\t%v\n", vv)
		}
	}

	fmt.Println("===============")
	sort.Sort(byAge(users))
	for _, v := range users {
		fmt.Printf("%v %v %v \n", v.First, v.Last, v.Age)
		for _, vv := range v.Sayings {
			fmt.Printf("\t%v\n", vv)
		}
	}
}
