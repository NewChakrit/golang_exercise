package main

import "fmt"

func main() {
	fmt.Println(paradise("Hawaii"))
}

func paradise(loc string) string {
	return fmt.Sprintf("My idea of paradise is %v", loc)
}
