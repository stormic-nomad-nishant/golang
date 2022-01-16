package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]   // Slice of all elements
	c := a[3:]  // Slice from 4rth element to end
	d := a[:6]  // Slice first 6 elements
	e := a[3:6] // Slice from 4rth to 6th element
	// a[5] = 23 // This will reflect in all the down below slices
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

}
