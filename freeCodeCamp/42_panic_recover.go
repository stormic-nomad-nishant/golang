package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end") // execution will continue
}
func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("ERROR:", err)
			// you can also panic here saying you can't handle it
			// panic(err) // If you do this, the call above the call stack will not coninue to end // this is equal to raising
		}
	}()
	panic("Something bad")
	fmt.Println("done panicking")
}
