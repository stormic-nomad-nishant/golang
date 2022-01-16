/* Different Types of Constants
- Named Constants
- Typed Constants
- Untyped Constants
- Enumerated constants
- Enumeration Expressions
*/
package main

import (
	"fmt"
)

func main() {
	// const MY_CONST // If you do this then this will be accessable even outside the package - Variable naming

	// TYPED CONSTANTS: you can declare constant values of any data types int, float32/64, string, bool
	const myNewConst int = 21
	fmt.Printf("%v,%T\n", myNewConst, myNewConst)
	// myNewConst = 45 // You cannot do this now to myNewConst as its a constant
	// const myConst float64 = math.Sin(1.45) // You cannot also do this as this will be determinded/computed at Compile time
	// This ^ means that you cannot set the value of a constant that is determined at RUN time
	const a = 21
	fmt.Printf("%v,%T\n", a, a) // You will notice the compiler will print the value of type as int
	// The compiler is inferring the value ^
	var b int16 = 27                // Differnt TYPE ^
	fmt.Printf("%v,%T\n", a+b, a+b) // This will NOT cause error by compiler because the compiler implecitly understands this way
}
