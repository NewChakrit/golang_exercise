package word

import (
	"fmt"
	"github.com/NewChakrit/golang_exercise/011-ninja-level-13/02/quote"
	"testing"
)

func TestCount(t *testing.T) {
	res := Count(quote.SunAlso)
	if res != 1349 {
		t.Error("Got: ", res, ", Expected: 1349")
	}
}

func ExampleCount() {
	fmt.Println(Count(quote.SunAlso))
	// Output:
	// 1349
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func TestUseCount(t *testing.T) {
	res := UseCount("one two theree")
	for k, v := range res {
		switch k {
		case "one":
			if v != 1 {
				t.Error("Got ", v, "Expected", 1)
			}
		case "two":
			if v != 1 {
				t.Error("Got ", v, "Expected", 2)
			}
		case "three":
			if v != 3 {
				t.Error("Got ", v, "Expected", 3)
			}
		}
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
