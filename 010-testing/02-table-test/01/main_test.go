package main

import (
	"testing"
)

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{[]int{21, 21}, 42},
		{[]int{1, 2, 3}, 6},
		{[]int{4, 7}, 11},
		{[]int{-1, 0, 1}, 0},
	}

	for _, v := range tests {
		y := mySum(v.data...)
		if y != v.answer {
			t.Errorf("Expected %v, Got %v", v.answer, y)
		}
	}

	x := mySum(2, 3)
	if x != 5 {
		t.Error("Expected", 5, "Got", x)
	}
}
