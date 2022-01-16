package main

import (
	"fmt"
)

const (
	_ = iota + 5 // _ is called the Write ONLY variable
	a
	b
	c
)

/*
_ is actually a write-only variable and you cannot get its value.
This is done because you must use all of the declared variables in the Go lang,
but sometimes you do not need to use all the return values you get from a function.
*/

func main() {
	// fmt.Printf("%v,%T\n", _, _) ==> YOU CANNOT READ THE VALUE OF THIS VARIABLE
	fmt.Printf("%v,%T\n", a, a)
}
