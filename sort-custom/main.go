package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.First, p.Age)
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

type ByName []Person

func (a ByName) Len() int {
	return len(a)
}

func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByName) Less(i, j int) bool {
	return a[i].First < a[j].First
}

func main() {
	p1 := Person{"James", 44}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)
	// [James: 44 Moneypenny: 27 Q: 64 M: 56]
	sort.Sort(ByAge(people))
	fmt.Println(people)
	// [Moneypenny: 27 James: 44 M: 56 Q: 64]

	sort.Sort(ByName(people))
	fmt.Println(people)
	// [James: 44 M: 56 Moneypenny: 27 Q: 64]
}
