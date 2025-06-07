package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()

	// FAN OUT
	// Distribute work across multiple functions (ten goroutines) that all read from in.
	xc := fanOut(in, 10) // This now returns []<-chan int, each channel producing factorials

	// FAN IN
	// multiplex multiple channels onto a single channel
	for n := range merge(xc...) { // merge() will now receive all channels from fanOut
		fmt.Println(n)
	}

	fmt.Println("Program finished.")
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j // Sends 100 numbers in total (10 * 10)
			}
		}
		close(out) // Close the channel after all numbers are sent
	}()
	return out
}

func fanOut(in <-chan int, n int) []<-chan int {
	var xc []<-chan int
	var wgFanOut sync.WaitGroup // <--- เพิ่ม WaitGroup สำหรับ fanOut

	for i := 0; i < n; i++ {
		wgFanOut.Add(1) // <--- เพิ่ม Go-routine ที่ต้องรอ
		// Each call to factorial creates a goroutine that reads from the 'in' channel
		// and writes factorials to its own 'out' channel.
		// Pass wgFanOut to factorial
		ch := factorial(in, &wgFanOut) // <--- ส่ง WaitGroup ไปให้ factorial
		xc = append(xc, ch)
	}

	// This goroutine waits for all factorial goroutines to finish and close their output channels.
	go func() {
		wgFanOut.Wait() // <--- รอให้ factorial goroutines ทั้งหมดเสร็จ
		// No need to close 'in' here, gen() already does it.
	}()

	return xc
}

// func factorial now accepts a WaitGroup pointer
func factorial(in <-chan int, wg *sync.WaitGroup) <-chan int { // <--- แก้ไข signature
	out := make(chan int)
	go func() {
		defer wg.Done() // <--- บอก WaitGroup เมื่อ Go-routine นี้เสร็จ
		for n := range in {
			out <- fact(n)
		}
		close(out) // Close 'out' after all numbers have been processed from 'in'.
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done. This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
