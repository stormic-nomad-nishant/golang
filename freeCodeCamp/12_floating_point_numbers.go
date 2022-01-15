package main

import (
	"fmt"
)

func main() {
	a := 4.5
	fmt.Printf("%v,%T\n", a, a)
	var n float32 = 22.43
	fmt.Printf("%v,%T\n", n, n)
	var k float64 = 13.7e72
	fmt.Printf("%v,%T\n", k, k)

}
