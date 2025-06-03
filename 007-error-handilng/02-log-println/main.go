package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happened", err) // err happened open no-file.txt: no such file or directory
		log.Println("err happened", err) // 2025/06/03 14:48:41 err happened open no-file.txt: no such file or directory
		//log.Fatalln(err)
		//panic(err)
	}
}

// Println formats using the default formats for its operands and writes to standard output.
