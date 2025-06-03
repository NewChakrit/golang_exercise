package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()

	_, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happened", err) // err happened open no-file.txt: no such file or directory
		//log.Println("err happened", err)
		log.Fatalln(err)
		//panic(err)
	}
}

func foo() {
	fmt.Println("When os.Exit is called, deferred function don't run")
}

//result
//2025/06/03 14:56:25 open no-file.txt: no such file or directory
//exit status 1
