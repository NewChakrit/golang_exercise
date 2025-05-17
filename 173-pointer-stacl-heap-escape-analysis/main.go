package main

import "fmt"

func main() {
	a := 1
	//fmt.Println(a)
	fmt.Println(&a)
}

/*
escape to the heap
variable shared between main() and Println()

moved to heap
variable moved to the heap
*/

// ===== command: go run -gcflags -m main.go =====

/* value : fmt.Println(a)
# command-line-arguments
./main.go:7:13: inlining call to fmt.Println
./main.go:7:13: ... argument does not escape
./main.go:7:14: a escapes to heap
1
*/

/* value : fmt.Println(&a)
# command-line-arguments
./main.go:8:13: inlining call to fmt.Println
./main.go:6:2: moved to heap: a
./main.go:8:13: ... argument does not escape
0x14000104020
*/
