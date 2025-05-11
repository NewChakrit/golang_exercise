package main

import (
	"bytes"
	"fmt"
)

//type person struct {
//	first string
//}
//
//func (p person) writeOut(w io.Writer) error {
//	_, err := w.Write([]byte(p.first))
//	return err
//}

func main() {
	//f, err := os.Create("output.txt")
	//if err != nil {
	//	log.Fatalf("error %s", err)
	//}
	//
	//defer f.Close()
	//
	//s := []byte("Hello Gophers!")
	//
	//_, err = f.Write(s)
	//if err != nil {
	//	log.Fatalf("error %s", err)
	//}

	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())
	b.WriteString("Gophers!")
	fmt.Println(b.String()) // Hello Gophers!
	b.Reset()
	b.WriteString("Hi dear ")
	fmt.Println(b.String()) // Hi dear

	b.Write([]byte("Happy Happy"))
	fmt.Println(b.String()) // Hi dear Happy Happy
}
