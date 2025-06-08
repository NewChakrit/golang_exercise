package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
}

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happened", err)
		//// result:
		//// err happened open no-file.txt: no such file or directory

		log.Println("err happened", err)
		// result:
		// 2025/06/08 15:13:23 err happened open no-file.txt: no such file or directory

		//log.Fatalln(err)
		//panic(err)
	}
}

/*
	Package log implements simple logging package ... write to
	standard error and prints the date and time of each
	logged message ... the Fatal functions call os.Exit(1)
    after writing the log message ... the Panic function call
	panic after writing the log message.
*/

// Println formats using the default formats for its
// operands and writes to standard output
