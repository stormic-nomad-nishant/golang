package main

import (
	"fmt"
)

func main() {
	a := [...]int{1, 2, 3, 4}
	b := a // This assignment means we are creating a brand new copy of array i.e go will create a new array b and copy the entire array a into b
	// Also this b:= a is copying the "VALUE OF THE ARRAY" and not address
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a) // The original array a is unchanged
	/*
	  You can use the & operator i.e POINTER (just like C/C++) to point to the address
	  instead of value as follows:
	*/
	fmt.Println("--POINTERS--")
	c := &a // array c is now pointing to array a
	c[1] = 5
	fmt.Println(a)
	fmt.Println(c)
}
