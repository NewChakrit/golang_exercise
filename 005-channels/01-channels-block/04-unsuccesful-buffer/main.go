package main

import "fmt"

func main() {
	c := make(chan int, 1) //  อนุญาติให้ค่าบางอย่างอยู่ในช่องโดยไม่สนว่ามีบางสิ่งพร้อมจะดึงออก

	c <- 42
	c <- 43

	fmt.Println(<-c) // fatal error: all goroutines are asleep - deadlock!
}
