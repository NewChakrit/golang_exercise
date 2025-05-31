package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)                // OS    darwin
	fmt.Println("ARCH\t\t", runtime.GOARCH)            // ARCH    arm64
	fmt.Println("CPUs\t\t", runtime.NumCPU())          // CPUs    10
	fmt.Println("Goroutine\t", runtime.NumGoroutine()) // Goroutine   1

	wg.Add(1)
	go foo()
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())          // CPUs    10
	fmt.Println("Goroutine\t", runtime.NumGoroutine()) // Goroutine   1
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
	wg.Done()
}
