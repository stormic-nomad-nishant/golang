package main

import (
	"fmt"
)

func main() {
	number := 50
	guess := 30
	if guess < number {
		fmt.Println("Too Low")
	}
	if guess > number {
		fmt.Println("Too high")
	}
	if guess == number {
		fmt.Println("You got it !!")
	}
	fmt.Println(number <= guess, number >= guess, number != guess)
	//
	//LOGICAL OPERATORS AND OR NOT
	if guess > 1 || guess < 100 {
		fmt.Println("Guess between 1 or 100")
	}
	if number > 30 && number < 70 {
		fmt.Println("Guess between 30 and 70")
	}
	fmt.Println(!true)

}
