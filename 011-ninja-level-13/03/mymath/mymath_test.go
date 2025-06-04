package mymath

import (
	"fmt"
	"testing"
)

var x = []int{1, 4, 6, 8, 100}

type test struct {
	data   []int
	answer float64
}

func TestCenteredAvg(t *testing.T) {
	tests := []test{
		{[]int{1, 4, 6, 8, 100}, 6},
		{[]int{0, 8, 10, 1000}, 9},
		{[]int{9000, 4, 10, 8, 6, 12}, 9},
		{[]int{123, 744, 140, 200}, 170},
	}

	for _, v := range tests {
		if CenteredAvg(v.data) != v.answer {
			t.Error("got", CenteredAvg(v.data), "expected:", v.answer)
		}
	}

	res := CenteredAvg(x)
	if res != 6 {

	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg(x))
	// Output:
	// 6
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg(x)
	}
}

// Command
// go test
// go test -cover
// go test -bench .
// go test -coverprofile c.out
// godoc -http=:8080
// go tool cover -html=c.out
