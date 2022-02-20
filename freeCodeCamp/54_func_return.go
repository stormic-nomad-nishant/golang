package main

import (
	"fmt"
)

func main() {
	returned_result := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is", returned_result)
}

func sum(values ...int) int { // The second int is the type of the return value, in this case a int
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result // return statement
}
