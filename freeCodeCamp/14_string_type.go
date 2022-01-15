package main

import (
	"fmt"
)

func main() {
	a := "This is a good example of string"
	fmt.Printf("%v,%T\n", a, a)
	var b string = "This is also a way to declare string"
	fmt.Printf("%v,%T\n", b, b)
	// Strings can be treated similer to arrays ex:
	fmt.Printf("%v, %T\n", string(b[2]), b[2])
	// Strings are generally immutable also
	fmt.Printf("%v\n", a+" --- "+b)
	// Converting a string into collection of bytes aka "Slice of bytes"
	c := []byte(a)
	fmt.Printf("%v,%T\n", c, c)
}
