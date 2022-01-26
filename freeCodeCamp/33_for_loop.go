package main

import (
	"fmt"
)

// NOTE: You don't have a comma (,) operator in golang
func main() {
	// Single loop
	fmt.Println("Single loop")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("Double loop")
	// Double value
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}
	fmt.Println("Missing intial condition from loop")
	// We can also do :
	k := 0 // Here k is only scoped to the for loop
	for ; k < 5; k++ {
		fmt.Println(k)
	}
	fmt.Println("Missing both intial condition & iteration from loop")
	// Other way Syntactic SUGAR
	p := 0
	for p < 5 {
		fmt.Println(p)
		p++
	}
	// Breaking out of a infinite loop with Breaking
	fmt.Println("The BREAK statement")
	h := 0
	for {
		fmt.Println(h)
		h++
		if h == 5 {
			break
		}
	}
	fmt.Println("The Continue statement")
	f := 0
	for ; f < 10; f++ {
		if f%2 == 0 {
			continue // i.e exit from the loop and start over
		}
		fmt.Println(f)
	}
}
