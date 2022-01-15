/*
Rune literals are just 32-bit integer values (however they're untyped constants,
so their type can change). They represent unicode codepoints.
For example, the rune literal 'a' is actually the number 97.
https://go.dev/doc/go1#rune

Simple term -> Rune is basically a int32
*/

package main

import (
	"fmt"
)

func main() {
	// implecit declare:
	r := 'a'                    // This is a rune and not a string because we are using a single quote and not the double quotes
	fmt.Printf("%v,%T\n", r, r) // Runes are just the type alias for int32 thats why you would see the type here as int32
	// explicit declare:
	var b rune = 'b'
	fmt.Printf("%v,%T\n", b, b)

}
