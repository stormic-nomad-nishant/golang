package main

import (
  "fmt"
)

func main() {
  var i int = 23
  fmt.Printf("%v,%T\n", i,i)
  var j float32
  // j = i  : This will cause a error because you cannot do a implicit type conversion, you have to tell the compiler as below and make it explicit
  j = float32(i) // Converting the int i to float32 j
  fmt.Printf("%v,%T\n", j,j)
}
