/*
A structure or struct in Golang is a user-defined type that allows to
group/combine items of possibly different types into a single type.

 Any real-world entity which has some set of properties/fields can be
 represented as a struct. This concept is generally compared with the classes
 in object-oriented programming. It can be termed as a lightweight class that
 does not support inheritance but supports composition.

 For Example, an address has a name, street, city, state, Pincode.
It makes sense to group these three properties into a single structure
address as shown below.

type Address struct {
      name string
      street string
      city string
      state string
      Pincode int
}


*/
package main

import (
	"fmt"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

// Structs can have any type of data types ^ i.e different data types unlike slices, maps
func main() {
	aDoctor := Doctor{
		number:    1,
		actorName: "Kreg",
		companions: []string{
			"Stone",
			"Meta",
			"Zeta",
		},
	}
	fmt.Println(aDoctor)
	// If u want to enquire one single value from struct then u can use the "."
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[2])
	// There is also another way to instantiate the structs in go with
	// positional syntax:
	bDoctor := Doctor{
		2,
		"Newton",
		[]string{
			"Alpha",
			"Beta",
			"Gamma",
		},
	}
	fmt.Println(bDoctor.actorName)
	// THOUGH it is advised not to use this syntex ^^ coz you will have to ensure
	// that you maintain the "positions" of the variables
}
