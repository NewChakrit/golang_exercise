package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("James.")
	if s != "Welcome my dear James." {
		t.Error("got", s, "expected:", "Welcome my dear James.")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("James."))
	// Output:
	// Welcome my dear James.
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James.")
	}
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
