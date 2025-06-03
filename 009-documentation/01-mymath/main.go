package main

// Sum adds an unlimited number of values of type int.
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}

	return sum
}

// cml
// go doc Sum

// go doc with agrument
// zero: go doc
// one: go doc <sym>[.<method>]
//      go doc [pkg.]<sym>[.<method>]
//      go doc [pkg.][<sym>].<method>
// two: go doc <sym>[.<method>]

//ex go doc json.number
