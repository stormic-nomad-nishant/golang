package main

import (
	"fmt"
)

func main() {
	func() { // This is an anonymous function . This is also a example of function as type
		fmt.Println("Hello go!!!")
	}() // This is how we call/invoke the anonymous func
}
