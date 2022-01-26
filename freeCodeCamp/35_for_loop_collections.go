package main

import (
	"fmt"
)

func main() {
	s := []int{11, 12, 13}
	// To access a slice or (array as well) ^ , you can use the range keyword as
	// below.. very similer to python

	for k, v := range s {
		fmt.Println(k, v)
	}
	// For maps (dicts) also u can try:
	statePopulation := map[string]int{
		"Delhi":    192,
		"UP":       214,
		"Mumbai":   34,
		"Banglore": 24,
	}
	for k, v := range statePopulation {
		fmt.Println("For the key :", k, "--> The value is:", v)
	}
	// for a string
	text := "Some importent text"
	for k, v := range text {
		fmt.Println(k, string(v))
	}
	// ALSO NOTE if you just need the keys.. you can skip declaring the value
	fmt.Println("Only printing the Keys")
	for k := range statePopulation {
		fmt.Println("For the key :", k)
	}
	// To pring only values, you can use the write only variable
	fmt.Println("Only values")
	for _, v := range statePopulation {
		fmt.Println("For the Values :", v)
	}
}
