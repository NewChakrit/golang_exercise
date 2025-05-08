package main

import "fmt"

func main() {
	//am := map[string]int{
	//	"Todd":   42,
	//	"Henry":  16,
	//	"Padget": 14,
	//}

	//fmt.Println(am)
	//fmt.Printf("%#v\n", am)

	//fmt.Println("------------")

	bm := make(map[string]int)
	bm["Lisa"] = 22
	bm["Bob"] = 35
	bm["George"] = 78

	//fmt.Println(bm)
	//fmt.Printf("%#v\n", bm)
	//fmt.Println(len(bm)) // 2

	// Delete
	delete(bm, "Lisa")
	delete(bm, "Tim") // Not thing change

	fmt.Println(bm)        // map[Bob:35 George:78]
	fmt.Println(bm["Tim"]) // 0

	tm := make(map[string]string)
	tm["Lisa"] = "Girl"
	tm["Bob"] = "Boy"
	tm["George"] = "Man"
	fmt.Println(tm)
	fmt.Println(tm["Tim"])      // ""
	fmt.Printf("%T", tm["Tim"]) // string

	rm := make(map[string]bool)
	rm["Lisa"] = true
	rm["Bob"] = false
	fmt.Println(rm)
	fmt.Println(rm["Tim"])      // false
	fmt.Printf("%T", rm["Tim"]) // bool

}
