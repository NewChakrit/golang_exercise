package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("begin CPU:", runtime.NumCPU())      // 10
	fmt.Println("begin gs:", runtime.NumGoroutine()) //1

	var wg sync.WaitGroup
	wg.Add(2)

	//fmt.Println(runtime.NumGoroutine())

	go func() {
		fmt.Println("Hello from A")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello from B")
		wg.Done()
	}()

	fmt.Println("mid CPU:", runtime.NumCPU())      // 10
	fmt.Println("mid gs:", runtime.NumGoroutine()) // 3

	wg.Wait()

	fmt.Println("about to exist.")
	fmt.Println("after CPU:", runtime.NumCPU())      // 10
	fmt.Println("after gs:", runtime.NumGoroutine()) //1
}
