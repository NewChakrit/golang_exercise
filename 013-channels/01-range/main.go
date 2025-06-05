package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c { // นี่คือ Go idiomatic way ในการรับค่าจาก Channel ที่ถูกออกแบบมาเพื่อจัดการกับการปิด Channel โดยอัตโนมัติ
		fmt.Println(n) // ลูปนี้จะ รับค่า n จาก Channel c ไปเรื่อยๆ ตราบใดที่ Channel ยังเปิดอยู่และมีข้อมูลอยู่
	}

}
