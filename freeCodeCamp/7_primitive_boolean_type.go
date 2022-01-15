/*
3 types of Primitive Data Types in GO:
  - Boolean type
  - Numeric type (int, float, complex number)
  - Text type
 */

// BOOLEAN DATA - TRUE or FALSE
package main

import (
  "fmt"
)


func main(){
  var n bool = true // or false
  fmt.Printf("%v,%T\n",n,n)
}
