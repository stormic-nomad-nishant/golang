/*
Arrays have their place, but they’re a bit inflexible, so you don’t see them too often in Go code.
Slices, though, are everywhere.
They build on arrays to provide great power and convenience.

The type specification for a slice is []T, where T is the type of the elements
of the slice.
Unlike an array type, a slice type has no specified length.
*/
package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3} // This is how you can initialize a slice without specifying the length
	fmt.Printf("%v\n", a)
	fmt.Println("Length: ", len(a)) // This gives you the length of slice
	// Slice also has a additional method called capacity.
	// As Slice is projection of array , so it can happen that the array that is backing
	// the slice might not be able to fully accomodate the slice
	fmt.Println("Capacity: ", cap(a))
	// Unlike Arrays, Slices are always by default 'reference types' i.e:
	b := a
	b[1] = 88
	fmt.Println(b)
	fmt.Println(a)
}
