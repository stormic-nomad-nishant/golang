package main

import (
	"fmt"
)

// NOTE: in golang, switch statement has implicit break statements, so u don't
// need to write it
func main() {
	fmt.Println("Switch Statements")
	// switch tag
	switch 2 {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Default case")
	}
	// You can also have multiple conditions in a case
	switch 24 {
	case 10, 12, 20:
		fmt.Println("TEN, TWELVE, TWENTY")
	case 21, 24, 45:
		fmt.Println("Second stuff")
	default:
		fmt.Println("Default case")
	}
	// You can also have this style:
	switch i := 1 + 2; i {
	case 1:
		fmt.Println("ONE")
	case 2:
		fmt.Println("TWO")
	default:
		fmt.Println("Default case as i is", i)
	}
	// One more way:
	k := 2
	switch {
	case k < 1:
		fmt.Println("k smaller than ONE")
	case k == 2:
		fmt.Println("TWO")
	default:
		fmt.Println("Default case as k is", k)
	}
}
