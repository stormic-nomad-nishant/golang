package main

import (
	"fmt"
)

func main() {
	var number int = 129
	if number > 1 && number < 20 {
		fmt.Println("Number is between 1 and 20")
	} else if number > 21 && number < 100 {
		fmt.Println("Number is between 21 and 100")
	} else {
		fmt.Println("Number is greater than 100")
	}

}
