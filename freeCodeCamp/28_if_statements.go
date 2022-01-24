package main

import (
	"fmt"
)

func main() {
	// BOOLEAN IF statement:
	// Also remember in Golang, you need to have the if statements wrapped with block
	// i.e { } else u get errors.
	if true {
		fmt.Println("This is inside if and should be printed")
	}
	// The other style : "Initializer Syntax"
	statePopulation := map[string]int{
		"Delhi":    192,
		"UP":       214,
		"Mumbai":   34,
		"Banglore": 24,
	}
	// This -->  pop, ok := statePopulation["Mumbai"] is initializer
	if pop, ok := statePopulation["Mumbai"]; ok {
		fmt.Println(pop)
	}
	// Also remember the block^ , as you cannot access the pop variable outside
	// the block : the below will be an error
	// fmt.Println(pop)

}
