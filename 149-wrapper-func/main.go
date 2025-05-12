package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	xb, err := readFile("poem.txt")
	if err != nil {
		log.Fatalf("readFile in main: %s", err)
	}

	fmt.Println(xb)
	fmt.Println(string(xb))
}

func readFile(fileName string) ([]byte, error) {
	xb, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error in readFile func: %s", err)
	}

	return xb, err
}
