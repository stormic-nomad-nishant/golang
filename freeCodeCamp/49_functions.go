// Basic syntax
package main

import (
	"fmt"
)

func main() {
	mycoolmssg("Hello Play")
}

func mycoolmssg(msg string) {
	fmt.Println("INSIDE the mycoolmssg method")
	fmt.Println(msg)
}
