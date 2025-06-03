package main

import (
	"fmt"
	"github.com/NewChakrit/golang_exercise/009-documentation/03-ninja-level-12/dog"
)

type canine struct {
	Name string
	Age  int
}

func main() {
	fido := canine{"Fido", dog.Years(10)}
	fmt.Println(fido)
}
