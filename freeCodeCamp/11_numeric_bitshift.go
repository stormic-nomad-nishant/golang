package main

import (
	"fmt"
)

func main() {
	a := 8              // 1000 in binary
	fmt.Println(a << 3) // Left shift a by 3 places
	fmt.Println(a >> 3) // Right shift a by 3 places
}
