package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}

// Terminal
// 1) go build acdc.go >> create main
// 2) ./main >> Hello World!
// 3) rm main >> main has been removed
// 4) go build . >> create 054...
// 5) ./054-hands-on-exercise-11 >> Hello World!
// 6) rm 054-hands-on-exercise-11   >> 054-hands-on-exercise-11  has been removed
// 7) go build ./... >> create 054...
