// TYPE CONVERSION
package main

import (
  "fmt"
  "strconv"
)

func main() {
  var i int = 42
  fmt.Printf("%v,%T\n",i,i)
  var j string
  /*
  Now in order to convert a int to string you cannot directly do :
  -> j = string(i)
  This will not work as expected..as number 42 will be converted to ascii code *
  So you will have to use the strconv package
  */
  j = strconv.Itoa(i) //Itoa -> Integer to ASCII (EXPLICIT)
  fmt.Printf("%v,%T\n",j,j)
}
