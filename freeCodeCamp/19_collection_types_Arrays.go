// Arrays and Slices are two type of Collection Typed in go
// 1. Arrays: can only store one type of data i.e homogenious
package main

import (
	"fmt"
)

func main() {
	// One way of declaring array
	grades := [4]int{1, 2, 3, 4}         // [size_of_array]datatypes{data in the array}
	var test [3]int = [3]int{12, 23, 45} // The other way of doing the same
	points := [3]float64{6.4, 2.1, 9.1}
	fmt.Println(grades)
	fmt.Println(test)
	fmt.Println(points)
	// Arrays are contigeous in memory & obviously Arrays starts with 0
	fmt.Println(points[0])
	fmt.Println(points[1])
	fmt.Println(points[2])
	// Other way : Use ... instead of defining the size of the array, this tells
	// the compiler to create a array just large enough to store my data
	new_grade := [...]int{67, 28, 321, 28, 91}
	fmt.Println(new_grade)
	var students [3]string // You can also declare a array as follows
	students[0] = "Nathan"
	students[1] = "Ram"
	students[2] = "tom"
	fmt.Println(students)
	fmt.Println(len(students)) // You can also get the length of the array with the len method
}
