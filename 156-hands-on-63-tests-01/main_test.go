package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(4, 3)
	if total != 7 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
