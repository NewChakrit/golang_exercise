package mystr

import (
	"strings"
	"testing"
)

const s = "We ask ourselves, Who am I to be brilliant, gorgeous, gorgeous, talented, fabulous? Actually, who are you not to be? Your playing small does not serve the world. There is nothing enlightened about shrinking so that other people won't fee insecure around you. We are all meant to shine, as children do. We were born to make manifest the glory that is within us."

func BenchmarkCat(b *testing.B) {
	xs := strings.Split(s, " ")

	for i := 0; i < b.N; i++ {
		Cat(xs)
	}

}

func BenchmarkJoin(b *testing.B) {
	xs := strings.Split(s, " ")

	for i := 0; i < b.N; i++ {
		Cat(xs)
	}
}
