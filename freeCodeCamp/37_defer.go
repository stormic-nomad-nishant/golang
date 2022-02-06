package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")
}

// The defered function are exicuted in LIFO order (Last in first out) i.e Stack
// Also remember the above defered functions are being executed after the main()
// method is done running but before it returns result to calling function
