package main

import (
	"fmt"
)

func main() {
	sum("The sum is", 1, 2, 3, 4, 5)
}

func sum(msg string, values ...int) {
	fmt.Println(values) // this acts as slice
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
}
