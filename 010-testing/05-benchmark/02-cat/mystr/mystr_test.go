package mystr

import (
	"fmt"
	"strings"
	"testing"
)

const s = "We ask ourselves, Who am I to be brilliant, gorgeous, gorgeous, talented, fabulous? Actually, who are you not to be? Your playing small does not serve the world. There is nothing enlightened about shrinking so that other people won't fee insecure around you. We are all meant to shine, as children do. We were born to make manifest the glory that is within us."

var xs = strings.Split(s, " ")

func TestCat(t *testing.T) {
	if s != Cat(xs) {
		t.Error("got", xs, "expected:", s)
	}
}

func BenchmarkCat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cat(xs)
	}
}

func ExampleCat() {
	fmt.Println(Cat(xs))
	// Output:
	// We ask ourselves, Who am I to be brilliant, gorgeous, gorgeous, talented, fabulous? Actually, who are you not to be? Your playing small does not serve the world. There is nothing enlightened about shrinking so that other people won't fee insecure around you. We are all meant to shine, as children do. We were born to make manifest the glory that is within us.
}

func TestJoin(t *testing.T) {
	if s != Join(xs) {
		t.Error("got", xs, "expected:", s)
	}
}

func BenchmarkJoin(b *testing.B) {
	xs := strings.Split(s, " ")

	for i := 0; i < b.N; i++ {
		Cat(xs)
	}
}

func ExampleJoin() {
	fmt.Println(Join(xs))
	// Output:
	// We ask ourselves, Who am I to be brilliant, gorgeous, gorgeous, talented, fabulous? Actually, who are you not to be? Your playing small does not serve the world. There is nothing enlightened about shrinking so that other people won't fee insecure around you. We are all meant to shine, as children do. We were born to make manifest the glory that is within us.
}

// Try cml:
// 1. go test -bench Greet || go test -bench . (to run all func)
// 2. go test -coverprofile <file name>
// EX go test -coverprofile c.out:
// ่คนี่คือ flag ที่สั่งให้ go test วัด Code Coverage และ บันทึกผลลัพธ์ลงในไฟล์
// coverprofile: เป็นชื่อ flag ที่บ่งบอกว่าเราต้องการสร้าง "coverage profile"
// c.out: คือชื่อไฟล์ที่คุณต้องการบันทึกผลลัพธ์ของ Code Coverage ลงไป คุณสามารถใช้ชื่อไฟล์อะไรก็ได้ที่คุณต้องการ (เช่น coverage.txt, coverage.out, etc.) แต่นิยมใช้ .out

// 3. คำสั่ง go tool cover -html=<file name>
//  เป็นคำสั่ง Go ที่ใช้ในการ สร้างรายงาน Code Coverage ในรูปแบบ HTML จากไฟล์ c.out ที่ได้จากการรัน Unit Tests ด้วย go test -coverprofile
