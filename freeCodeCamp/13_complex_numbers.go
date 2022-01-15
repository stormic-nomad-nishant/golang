package main

import (
	"fmt"
)

// 2 types of complex number: complex 64 and complex 128
func main() {
	var n complex64 = 1 + 24i // (complex number = real + imaginary part)
	fmt.Printf("%v,%T\n", n, n)
	var m complex128 = 1 + 24i
	fmt.Printf("%v,%T\n", m, m)
	fmt.Printf("%v,%T\n", real(m), real(m))
	fmt.Printf("%v,%T\n", imag(m), imag(m))
	// How to change a number to a complex number?
	var u complex128 = complex(5, 12) // complex(real,complex)
	fmt.Printf("%v,%T\n", u, u)

}
