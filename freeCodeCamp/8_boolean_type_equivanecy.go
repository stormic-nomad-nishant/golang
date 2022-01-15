package main

import (
    "fmt"
)
// NOTE: In general every time you create a variable , its value is 0 i.e in boolean its false 
func main(){
  a := 1 == 1 // Double == is called the EQUALS OPERATOR, here we are checking if item on left == right
  b := 1 == 2
  fmt.Printf("%v,%T\n",a,a)
  fmt.Printf("%v,%T\n",b,b)
}
