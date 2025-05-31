package main

import "fmt"

func main() {
	c := make(chan int, 2) //  อนุญาติให้ค่าบางอย่างอยู่ในช่องโดยไม่สนว่ามีบางสิ่งพร้อมจะดึงออก

	c <- 42
	c <- 43

	fmt.Println(<-c) // 42
	fmt.Println(<-c) // 43

	fmt.Println("-------------")
	fmt.Printf("%T\n", c) // chan int

}
