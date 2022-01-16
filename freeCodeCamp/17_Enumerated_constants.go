package main

import (
	"fmt"
)

// Enumerated constants
// iota is a counter that is used to create a enumerated constants
// const a = iota

const (
	a = iota
	b = iota
	c = iota
)

// Iota's are also scoped to the block ^
// The below const group will be inferred by the compiler and the compiler will know the pattern
// also this const group/block will again be initialized from 0 as this is differnt block from the above
const (
	alpha = iota
	beta
	charlie
)

func main() {
	fmt.Printf("%v,%T\n", a, a) // this will print 0 , intial value of iota = 0 (same as int)
	fmt.Printf("-------------\n")
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	// You will see Iota is incrementing its value for the constants ^^
	// Lets print the other const group : (iota is scoped to a constant block)
	fmt.Printf("**************\n")
	fmt.Printf("%v\n", alpha)
	fmt.Printf("%v\n", beta)
	fmt.Printf("%v\n", charlie)
}
