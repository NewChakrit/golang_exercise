package mystr

import "strings"

func Cat(xs []string) string {
	//s := ""
	s := xs[0]
	//for _, v := range xs {
	//s += v
	//s += " "

	for _, v := range xs[1:] {
		s += " "
		s += v
	}
	return s
}

func Join(xs []string) string {
	return strings.Join(xs, " ")
}
