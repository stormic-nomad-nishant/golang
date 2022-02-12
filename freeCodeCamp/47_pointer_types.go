package main

import (
	"fmt"
)

func main() {
	var ms *myStruct        // pointer to myStruct
	ms = &myStruct{foo: 49} // initializing the struct, but with address of operator
	fmt.Println(ms)
}

type myStruct struct {
	foo int
}
