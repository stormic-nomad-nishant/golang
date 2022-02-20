package main

import (
	"fmt"
)

func main() {
	returned_result := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is", returned_result)
}

// This approach is called "named return"
func sum(values ...int) (result int) { // Here we are declaring both the return variable and the type of it
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return // this returns the 'result' variable
}
