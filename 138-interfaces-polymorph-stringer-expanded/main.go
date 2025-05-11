package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("This is the book ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("The number is ", strconv.Itoa(int(c)))
}

func logInfo(s fmt.Stringer) {
	log.Println("LOG FROM 138", s.String())
}

func main() {
	b := book{
		title: "West with the night.",
	}

	var c count = 42

	//log.Println(b) //2025/05/11 22:56:41 This is the book West with the night.
	//log.Println(c) //2025/05/11 22:56:41 The number is 42

	logInfo(b) //2025/05/11 23:00:16 LOG FROM 138 This is the book West with the night.
	logInfo(c) //2025/05/11 23:00:16 LOG FROM 138 The number is 42

}
