package main

import (
	"fmt"
	"log"
)

/*
Go language provides a special feature known as an anonymous function.
An anonymous function is a function which doesn’t contain any name.
It is useful when you want to create an inline function.
In Go language, an anonymous function can form a closure.
An anonymous function is also known as function literal.
*/
func main() {
	fmt.Println("Start")
	defer func() { // This method is called as anonymouns function
		if err := recover(); err != nil {
			log.Println("IN RECOVER - Error -->", err)
		}
	}() // This is how we invoke an anonymous function
	panic("SOMETHING BAD HAPPENED")
	fmt.Println("END")

}
