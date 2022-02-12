package main

import (
	"fmt"
)

func main() {
	var a int = 45
	var b *int = &a
	fmt.Println(a, *b)
	a = 24 // If you change a here, the value of b would also change as b is pointing to a
	fmt.Println(a, *b)
	*b = 55 // If you change the value of b i.e this means that you are now changing a value
	fmt.Println(a, *b)
}
