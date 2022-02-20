package main

import (
	"fmt"
)

func main() {
	f := func() { // This is an anonymous function . This is also a example of function as type
		fmt.Println("Hello go!!!")
	}
	f() // This is how we call/invoke the anonymous func
	// OR
	var n func() = func() {
		fmt.Println("Hello new go!!!!!")
	}
	n()

}
