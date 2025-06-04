package main

import (
	"fmt"
	"github.com/NewChakrit/golang_exercise/010-testing/05-benchmark/02-cat/mystr"
	"strings"
)

const s = "We ask ourselves, Who am I to be brilliant, gorgeous, gorgeous, talented, fabulous? Actually, who are you not to be? Your playing small does not serve the world. There is nothing enlightened about shrinking so that other people won't fee insecure around you. We are all meant to shine, as children do. We were born to make manifest the glory that is within us."

func main() {
	xs := strings.Split(s, " ")

	for _, v := range xs {
		fmt.Println(v)
	}

	fmt.Printf("\n%s\n", mystr.Cat(xs))
	fmt.Printf("\n%s\n\n", mystr.Join(xs))
}
