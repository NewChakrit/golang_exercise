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

// ===== command: go run -gcflags -m acdc.go =====

/* value : fmt.Println(a)
# command-line-arguments
./acdc.go:7:13: inlining call to fmt.Println
./acdc.go:7:13: ... argument does not escape
./acdc.go:7:14: a escapes to heap
1
*/

/* value : fmt.Println(&a)
# command-line-arguments
./acdc.go:8:13: inlining call to fmt.Println
./acdc.go:6:2: moved to heap: a
./acdc.go:8:13: ... argument does not escape
0x14000104020
*/
