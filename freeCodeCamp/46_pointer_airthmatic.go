package main

import (
	"fmt"
)

func main() {
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	fmt.Printf("%v %p %p", a, b, c)
	// c := &a[1] - 4 // you cannot do such sort of airthmatic opertions in go because
	// of the overall complexity involved, hence go doesn't allow pointer airthmatic
}
