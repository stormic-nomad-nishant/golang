// Basic syntax
package main

import (
	"fmt"
)

func main() {
	mycoolmssg("Heyyy", "Nishant")
}

// In this below syntax , the go compiler will understand  that all the arguments
// are of type string . So this is one way of passing the argument
func mycoolmssg(msg, name string) {
	fmt.Println(msg, name)
}
