package main

import "fmt"

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is error: %v", ce.info)
}

func main() {
	c1 := customErr{"need more coffee"}
	foo(c1)
}

func foo(e error) {
	fmt.Println("foo ran -", e, "\n", e.(customErr).info)
}

//result
//foo ran - here is error: need more coffee
//need more coffee
