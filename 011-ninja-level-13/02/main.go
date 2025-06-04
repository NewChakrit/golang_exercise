package main

import (
	"fmt"
	"github.com/NewChakrit/golang_exercise/011-ninja-level-13/02/quote"
	"github.com/NewChakrit/golang_exercise/011-ninja-level-13/02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso)) // 1349

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
