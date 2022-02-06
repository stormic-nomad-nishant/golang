package main

import (
	"fmt"
)

// Go doesn't have exceptions (try/catch in python) because unlike other languages
// go has normal response to exceptions. Like if you try opening file and it does
// not exist. But there are some situations which will push a go code into a state
// where it cannot continue , go offers "PANIC" to handle such cases.
func main() {
	fmt.Println("Start")
	panic("SOMETHING BAD IS HAPPENING")
	fmt.Println("END")
}
