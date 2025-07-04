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

	// ปัญหา: หากผู้ส่ง (done <- true) พยายามส่งสัญญาณ
	//แต่ไม่มีผู้รับ (<-done) ที่พร้อมจะรับ ทันที ผู้ส่งจะ บล็อก (block)

	// we block here until done <- true
	<-done
	<-done
	close(c)

	// to unblock above
	// we need to take values off of chan c here
	// but we never get here, because we're blocked above

	for n := range c {
		fmt.Println(n)
	}

}
