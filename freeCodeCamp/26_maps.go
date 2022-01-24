/*
Maps are a powerful, ingenious, and versatile data structure.
Golang Maps is a collection of unordered pairs of key-value.
It is widely used because it provides fast lookups and values that can retrieve,
update or delete with the help of keys.
- Maps are equivalent to the Dicts in python
- It is also a reference of hash table
*/
package main

import (
	"fmt"
)

func main() {
	// map[KEY]value
	animals := map[int]string{
		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}
	// NOTE: A map "type" has to be always consistent with the type of KEY:VALUE pair
	// Ex in above we have [int] as "Keys" and Strings as "values"
	statePopulation := map[string]int{
		"Delhi":    192,
		"UP":       214,
		"Mumbai":   34,
		"Banglore": 24,
	}
	fmt.Println(animals[93])
	fmt.Println(statePopulation["Banglore"])

	// MAKE fucntion can also be used to create maps
	CountryPopulation := make(map[string]int)
	CountryPopulation = map[string]int{
		"INDIA": 19241234,
		"CHINA": 1282138,
		"JAPAN": 124,
	}
	fmt.Println(CountryPopulation["INDIA"])
	fmt.Println(CountryPopulation)

	//ADDING A KEY
	CountryPopulation["USA"] = 1241241241414
	fmt.Println(CountryPopulation)

	// ALSO NOTE: Return order is NOT GURANETEED

	// DELETING a KEY
	delete(CountryPopulation, "CHINA")
	fmt.Println(CountryPopulation)
	// Interstingly if you again check for a key in general in a map which is not
	// present, it doesn't return in error but rather 0
	fmt.Println(CountryPopulation["HAHAH"])
	// or
	fmt.Println(CountryPopulation["CHINA"])
	// So this is VERY confusing coz it can lead you to falsly beleive that if the
	// value is 0 or not found. so in order to solve this, you inquire the dict like
	// below:
	data, flag := CountryPopulation["CHINA"]
	fmt.Println(data, flag)
	// This will give u 0, false ^ , and similery:
	data1, flag2 := CountryPopulation["INDIA"]
	fmt.Println(data1, flag2)
	// This will give u the value for "INDIA" and the flag as true indicating that its there
	// REMEMBER you cannot reuse the same variable
	// ALSO you use the WRITE Only variable to print the flag
	_, flag3 := CountryPopulation["JAPAN"]
	fmt.Println(flag3)
	// FINDING number of elements in a map
	fmt.Println("Length of the CountryPopulation MAP : ", len(CountryPopulation))
	// Just like slices, if you have multiple variables which are pointing to a map,
	// then the map is also passed by Reference , for example in below code:
	secound_world_coutries := CountryPopulation
	delete(secound_world_coutries, "JAPAN")
	fmt.Println("LETS SEE WHAT HAPPENS TO CountryPopulation now", secound_world_coutries)
}
