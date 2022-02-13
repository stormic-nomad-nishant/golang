// Basic syntax
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		mycoolmssg("Heyyy this is cool", i)
	}
}

func mycoolmssg(msg string, indx int) {
	fmt.Println(msg)
	fmt.Println("The value of index is ", indx)
}
