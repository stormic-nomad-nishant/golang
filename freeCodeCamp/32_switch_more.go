package main

import (
	"fmt"
)

func main() {
	i := 10
	switch {
	case i <= 10:
		fmt.Println("Less than or equal to 10")
		// this would print ^ but ideally the below condition is also true, so you
		// can use 'fallthrough'
		fallthrough // also know fallthrough is logic less and its upto u to use it correctly, coz it doesn't check for the condition in this case if case i <=20 or if case >= 20
		// So fallthrough should be only used when you want to run the next case as well
	case i <= 20:
		fmt.Println("Less than or equal to 20")
	default:
		fmt.Println("Default case")
	}
}
