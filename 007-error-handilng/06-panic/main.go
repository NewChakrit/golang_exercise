package main

import (
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		panic(err)
	}
}

// panic: open no-file.txt: no such file or directory
//
//goroutine 1 [running]:
//main.main()
//        /Users/a657002/Desktop/Learn/udemy/howtolearngo/golang_exercise/007-error-handilng/06-panic/main.go:10 +0x50
//exit status 2
