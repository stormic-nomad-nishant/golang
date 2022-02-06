package main

import (
	"fmt"
)
// Panics happen after defer statement
func main() {
	fmt.Println("start")                             // 1st this is being first printed
	defer fmt.Println("This statement was deferred") // 4rth : This is printed
	panic("some bad stuff happened")                 // 2nd compiler sees this
	fmt.Println("End")
}

// 3. compiler exits main and looks for defered
