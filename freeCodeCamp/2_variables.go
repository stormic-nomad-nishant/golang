/*
There are three different ways to declare the variable at function level:
1. Declare the variable itself. ex: var name_of_variable type_of_variable
2. Combine the declare and initialize it at same time
3. Let go decide the type of variable on its own
*/
package main // Packages are how code is organized in sub liberaries in go.

import "fmt" // importing a fmt package

func main() {
	var age int // 1st way: This is one way of declaring variable
	age = 20
	var age2 int = 45 // 2nd way: Declare a variable and initialize
	age3 := 89        // 3rd way: Go compiler can figure out the type of variable on its own
	name := "nishant"
	fmt.Println(age)
	fmt.Println(age2)
	fmt.Println(age3)
	fmt.Println(name)
	// Println formats using the default formats for its operands and writes to standard output. Spaces are always added between operands and a newline is appended. It returns the number of bytes written and any write error encountered.
	// Printf formats according to a format specifier and writes to standard output. It returns the number of bytes written and any write error encountered.
	// To print a variable’s type, you can use the %T (T means type) verb in the fmt.Printf function format. It’s the simplest and most recommended way of printing type of a variable.
	fmt.Printf("age3 is of type: %T\n", age3)
	fmt.Println("My name is: ", name)
	fmt.Printf("Type of name variable is: %T\n", name)
	// Alternatively, you can use TypeOf function from the reflection package reflect. However, it uses complex and expensive runtime reflection
	// fmt.Printf("Type of name variable is: %s\n", reflect.TypeOf(name))
	fmt.Printf("%v,  %T\n", name, name) // %v prints value of varible
}
