package main

import (
    "fmt"
)

// As mention previously , In general every time you create a variable it gets a value 0
// INT - size usually is 32 bits


func main(){
    // Signed bit integer range : A signed integer is a 32-bit datum that encodes an integer in the range [-2147483648 to 2147483647]
    // 8 bit integer range --> int8 : -128 to 127
    // 16 bit integer range --> int16: -32768 to 32767
    // 32 bit integer range --> int32: -2147483648 to 2147483647
    // 64 bit integer range --> int64: -9quintalian to 9quintalian
    var n int = 12 // SIGNED INTEGER
    fmt.Printf("%v,%T\n",n,n)
    // UNsigned bit integer range : An unsigned integer is a 32-bit datum that encodes a nonnegative integer in the range [0 to 4294967295].
    // uint8 : 0-255
    // uint16: 0-65535
    // uint32: 0-4294967295
    var m uint16 = 12
    fmt.Printf("%v,%T\n",m,m)
    // An important thing to note:
    var a int = 5
    var b int = 10
    fmt.Printf("%v\n",a+b)
    // The above would work , but the below:
    var c int8 = 20

    // fmt.Printf("%v\n",a+c)
    // Will cause an error: ^ ./9_numeric_int_type.go:31:24: invalid operation: a + c (mismatched types int and int8)
    // So to solve you have to do a "explicit" type conversion
    fmt.Printf("%v\n",a+int(c))
}
