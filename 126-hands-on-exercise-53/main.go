package main

import "fmt"

type person struct {
	firstName        string
	lastName         string
	favoriteIcecream []string
}

func main() {
	a := person{
		firstName:        "Name",
		lastName:         "Licha",
		favoriteIcecream: []string{"mango", "vanilla"},
	}

	b := person{
		firstName:        "Bob",
		lastName:         "Maya",
		favoriteIcecream: []string{"chocolate", "strawberry", "durian"},
	}

	fmt.Printf("%v", a)
	fmt.Printf("%v\n", a.firstName)
	fmt.Printf("%v\n", a.lastName)
	fmt.Printf("%v\n", a.favoriteIcecream)

	fmt.Printf("%v", b)
	fmt.Printf("%v\n", b.firstName)
	fmt.Printf("%v\n", b.lastName)
	fmt.Printf("%v\n", b.favoriteIcecream)

	fmt.Printf("Most faverite ice-cream of %v  %v\n", a.firstName, a.lastName)
	for i, v := range a.favoriteIcecream {
		i++
		fmt.Printf("%v. %v \n", i, v)
	}

	fmt.Println("---------------")

	fmt.Printf("Most faverite ice-cream of %v  %v\n", b.firstName, b.lastName)
	for i, v := range b.favoriteIcecream {
		i++
		fmt.Printf("%v. %v \n", i, v)
	}
}
