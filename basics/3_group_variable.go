// Variable at Package level
// You will have to use the full declaration syntex
package main

import (
	"fmt"
)

var i float64 = 89.23 // Declaration at package level

/*
You can also use a variable block to declare a group of Variable
by wrapping the variables in a var block as follows:
*/
var (
	name string  = "nishant"
	age  int     = 30
	pay  float32 = 4.4
)

func main() {
	fmt.Printf("%v,%T\n", i, i)
	fmt.Printf("Hi my name is %v, my age is %v and my salary is %v\n", name, age, pay)
}
