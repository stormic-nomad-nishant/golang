package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	// removing element from begining of slice
	b := a[1:] // This creates a new slice and removes the 1st number of a
	fmt.Println(b)
	// Removing element from end of slice
	c := []int{1, 2, 3, 4, 5}
	d := c[:len(c)-1]
	fmt.Println(d)
	// Removing elemet from middle of slice
	e := []int{1, 2, 3, 4, 5, 6}
	f := append(e[:2], a[3:]...)
	fmt.Println(f) // weird way ^
}
