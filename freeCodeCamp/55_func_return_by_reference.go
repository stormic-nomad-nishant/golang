package main

import (
	"fmt"
)

func main() {
	returned_result := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is", *returned_result) // We are accessing the result
}

// GO can let u return a local variabke as a pointer
func sum(values ...int) *int { // The second int is the type of the return value, in this case a pointer to int
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result // return statement here contains address of the local result . This value is written on the shared memory i.e heap
}
