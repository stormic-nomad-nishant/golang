package main

import (
	"fmt"
)

func main() {
	a := 42
	b := a // go will just copy the data from variable a to variable b
	fmt.Println(a, b)
	a = 27
	fmt.Println(a, b)
	// This is call by value ^
	fmt.Println("-------------")
	var c int = 33
	var d *int = &c // & is called the address of operator. // The * here before the type int is declaring a pointer to data of that type. i.e pointer to an integer
	// Above we diclared d as a pointer and assigned the value of d to point to
	// the address of c
	fmt.Println(&c, d) // Here you can see the address of c and d holding the address of c
	fmt.Println(c, *d) // Here * is called the dereferencing operator to print the value of d which it is holding
}
