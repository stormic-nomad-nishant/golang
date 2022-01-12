package main

import (
	"fmt"
)

// There are only 3 level of visibility of variable in GoLang
// 1. At Package Level:
var (
	i int = 20 // Package level declaration of i, also note this is lower case which means it is scoped to the package only
)
// Lowercase variable name means the variable is scoped to the package only i.e any file in the same package can access this variable.
// But if you do the following:
// 2. At this : This variable is exposed outside the package level & it is GLOBALLY VISIBLE
var I int = 92 // Note this is a upper_case variable and this is called a GLOBAL VARIABLE

// 3. At function block level: The below varible is at function level or block level hence called blocked scope, this below i cannot be accessed outside the block (i.e cann't be accessed outside main)
func main() {
	fmt.Printf("%v\n",i)
	var i int = 40        // Shadow variable as i is declared in the fucntion as well
	fmt.Printf("%v\n", i) // This will print the value as 40 because the function level variable  takes more precedence
}

// Length of a variable name also depicts the life span of a variable name
