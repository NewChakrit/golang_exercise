package main

import "fmt"

// func (r receiver) identifier(p parameter(s)) (return(s)) { code }

/*
func syntax

no params, no returns
1 param, no returns
1 param, 1 return
2 params , 2 returns

*/

func main() {
	foo()
	bar("New")

	fmt.Println(aloha("May"))

	owner, ageDogYears := dogYears("Todd", 4)
	fmt.Printf("%v : %v\n", owner, ageDogYears)
}

func foo() {
	fmt.Println("I'm foo.")
}

func bar(s string) {
	fmt.Println("My name's", s, ".")
}

func aloha(s string) string {
	return fmt.Sprintf("My name's %v.", s)
	//return fmt.Sprint("My name's ", s, ".")
}

func dogYears(name string, age int) (string, int) {
	age *= 7 // age = age * 7

	return fmt.Sprint(name, " is this old in dog years"), age
}
