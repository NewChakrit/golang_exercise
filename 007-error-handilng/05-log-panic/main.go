package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Panicf("err happened", err) // err happened open no-file.txt: no such file or directory
		//log.Println("err happened", err)
		//log.Fatalln(err)
		//panic(err)
	}
}

// 2025/06/03 15:01:22 err happened%!(EXTRA *fs.PathError=open no-file.txt: no such file or directory)
//panic: err happened%!(EXTRA *fs.PathError=open no-file.txt: no such file or directory)
//
//goroutine 1 [running]:
//log.Panicf({0x100d31b63?, 0xb?}, {0x1400007cf28?, 0x0?, 0x1400007cf38?})
//        /opt/homebrew/opt/go/libexec/src/log/log.go:439 +0x64
//main.main()
//        /Users/a657002/Desktop/Learn/udemy/howtolearngo/golang_exercise/007-error-handilng/05-log-panic/main.go:11 +0x68
//exit status 2
