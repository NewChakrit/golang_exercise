package main

import "fmt"

type engine struct {
	electric bool
}

func main() {

	a := struct {
		engine engine
		make   int
		model  string
		doors  int
		color  string
	}{
		make:   1996,
		model:  "Corrolado",
		doors:  4,
		color:  "black",
		engine: engine{electric: false},
	}

	b := struct {
		engine engine
		make   int
		model  string
		doors  int
		color  string
	}{
		make:   2018,
		model:  "Camry",
		doors:  4,
		color:  "Gray",
		engine: engine{electric: true},
	}

	fmt.Println(a)
	fmt.Println(b)

	e := map[string]bool{
		"electric A": a.engine.electric,
		"electric B": b.engine.electric,
	}

	fmt.Println("------------------")

	fmt.Printf("Model of vahicle A is %v\n", a.model)
	fmt.Printf("Vahicle A made in %v.\n", a.make)
	fmt.Printf("Vahicle A have %v doors.\n", a.doors)
	fmt.Printf("Color of vahicle A is %v.\n", a.color)
	for k, v := range e {
		if k == "electric B" {
			break
		}
		fmt.Printf("%v: %v\n", k, v)
	}

	fmt.Printf("Model of vahicle A is %v\n", a.model)
	fmt.Printf("Vahicle A made in %v.\n", a.make)
	fmt.Printf("Vahicle A have %v doors.\n", a.doors)
	fmt.Printf("Color of vahicle A is %v.\n", a.color)
	for k, v := range e {
		if k == "electric A" {
			continue
		}
		fmt.Printf("%v: %v\n", k, v)
	}

}
