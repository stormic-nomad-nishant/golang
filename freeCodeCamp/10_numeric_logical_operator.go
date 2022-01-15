//BIT OPERATIONS
package main

import (
    "fmt"
)

func main(){
  a := 10 // This is interpretted as binary 1010
  b := 3  // This is interpretted as binary 0011
  fmt.Println(a&b) // logical AND  - 0010
  fmt.Println(a|b) // logical OR   - 1011
  fmt.Println(a^b) // exclusive OR - either one number has the bit set or the other number but NOT both number - 1001
  fmt.Println(a&^b) // AND NOT - Only if neither of the number bit is set 0100
}
