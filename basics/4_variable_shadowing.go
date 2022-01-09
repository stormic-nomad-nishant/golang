package main

import (
	"fmt"
)

var (
	i int = 20 // Package level declaration of i
)

func main() {
	var i int = 40        // Shadow variable as i is declared in the fucntion as well
	fmt.Printf("%v\n", i) // This will print the value as 40 because the function level variable  takes more precedence
}
