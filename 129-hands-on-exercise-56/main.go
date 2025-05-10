package main

import "fmt"

func main() {

	a := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "New",
		friends: map[string]int{
			"Mark":  17,
			"Carol": 16,
			"Penny": 18,
		},
		favDrinks: []string{"Coke", "Est cola", "Mile", "Soy milk"},
	}

	fmt.Println("My name is", a.first)
	fmt.Println("These are list of my friends and their age.")
	for k, v := range a.friends {
		pronoun := "She"

		if k == "Mark" {
			pronoun = "He"
		}

		fmt.Printf("%v, %v's %v years old.\n", k, pronoun, v)
	}

	fmt.Println("-------------------")

	fmt.Println("These are list of ma fav drink.")
	for i, v := range a.favDrinks {
		i++
		fmt.Printf("%v. %v\n", i, v)
	}

}
