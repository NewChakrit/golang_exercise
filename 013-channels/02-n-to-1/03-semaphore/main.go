package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(c)
	}()

	for n := range c { // ลูปนี้จะ รับค่า n จาก Channel c ไปเรื่อยๆ ตราบใดที่ Channel ยังเปิดอยู่และมีข้อมูลอยู่
		fmt.Println(n)
	}

}

// ในโค้ดนี้ Producer แต่ละตัว "ปล่อย" สัญญาณไปยัง done หนึ่งครั้ง
//และ Coordinator "รอ" รับสัญญาณครบจำนวนที่กำหนด (2 ครั้ง)
//ก่อนที่จะดำเนินการปิด Channel c นี่คือพฤติกรรมที่คล้ายกับ Barrier
//หรือการใช้ Semaphore เพื่อนับจำนวนงานที่เสร็จสิ้น
