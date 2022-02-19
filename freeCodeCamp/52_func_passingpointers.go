package main

import (
	"fmt"
)

func main() {
	greeting := "Hello"
	name := "Nishant"
	sayGreeting(&greeting, &name)
	fmt.Println(name)
}

func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Mahek"
	fmt.Println(*name)
}
