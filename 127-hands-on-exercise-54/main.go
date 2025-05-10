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
		favoriteIcecream: []string{"mango", "vanilla", "chocolate"},
	}

	b := person{
		firstName:        "Bob",
		lastName:         "Maya",
		favoriteIcecream: []string{"chocolate", "strawberry", "durian"},
	}

	c := map[string]person{
		a.lastName: a,
		b.lastName: b,
	}

	fmt.Println(c)

	fmt.Println("---------------------")

	for k, v := range c {
		fmt.Printf("Key:%v \n", k)
		fmt.Printf("%v \n", v)
		fmt.Println("First name: ", v.firstName)
		fmt.Println("Last name: ", v.lastName)
		fmt.Println("Love fav ice-cream: ")
		for i, v2 := range v.favoriteIcecream {
			i++
			fmt.Printf("%v. %v\n", i, v2)
		}

		fmt.Println("---------------------")
	}
}
