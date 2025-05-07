package main

import "fmt"

func main() {
	jb := []string{"Jame", "Bond", "Martini", "Chocolate"}
	jm := []string{"Jenny", "Moneypenny", "Guiness", "Wolverine Tracks"}

	fmt.Println(jb)
	fmt.Println(jm)

	xxs := [][]string{jb, jm}
	fmt.Println(xxs)
	// [[Jame Bond Martini Chocolate] [Jenny Moneypenny Guiness Wolverine Tracks]]

}
