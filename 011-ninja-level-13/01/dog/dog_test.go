package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	s := Years(10)
	if s != 70 {
		t.Error("Got:", Years(10), ", Expected: ", "70")
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	// Output:
	// 70
}

func BenchmarkYear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func TestYearsTwo(t *testing.T) {
	s := YearsTwo(20)
	if s != 140 {
		t.Error("Got:", YearsTwo(20), ", Expected: ", 140)
	}
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(20))
	// Output:
	// 140
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(20)
	}
}
