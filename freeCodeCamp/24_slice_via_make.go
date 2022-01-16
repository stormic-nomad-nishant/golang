package main

import (
	"fmt"
)

// You can also use the make statement to create the slice
func main() {
	a := make([]int, 3) // 2 arguments (the type of slice, the length of slice)
	fmt.Printf("%v\n", a)
	fmt.Println("Length: ", len(a))
	fmt.Println("Capacity: ", cap(a))
	b := make([]int, 3, 100) // 3 arguments (the type of slice, the length of slice, the capacity of slice)
	fmt.Printf("%v\n", b)
	fmt.Println("Length: ", len(b))
	fmt.Println("Capacity: ", cap(b))
	// Unlike arrays, slices don't have to have a fixed size throughout its life
	// You can add and delete elements off slices
	fmt.Println("Lets see how to add/delete elements of slice :")
	new_slice := []int{} // a empty slice (python equivalent of list)
	fmt.Println(new_slice)
	fmt.Println("Length of new_slice BEFORE APPEND: ", len(new_slice))
	fmt.Println("Capacity of new_slice BEFORE APPEND: ", cap(new_slice))
	new_slice = append(new_slice, 1) // Append syntax to slice // ALSO note the 2nd 'new_slice' i.e the one being assigned to is a new slice that go would create and copy the old one to it
	// Append can take 2 or more arguments
	fmt.Println(new_slice)
	fmt.Println("Length of new_slice AFTER APPEND: ", len(new_slice))
	fmt.Println("Capacity of new_slice AFTER APPEND: ", cap(new_slice))
	new_slice = append(new_slice, 2, 3, 4, 5, 6, 7, 8) // The append function can take 2 or more arguments
	fmt.Println("After 2nd Append")
	fmt.Println(new_slice)
	fmt.Println("Length of new_slice 2nd AFTER APPEND: ", len(new_slice))
	fmt.Println("Capacity of new_slice 2nd AFTER APPEND: ", cap(new_slice))
	// CONCATINATE 2 SLICES:
	new_slice = append(new_slice, []int{67, 77, 99}...) // You have to add the last ...
	fmt.Println("Concatination of 2 slices:")
	fmt.Println(new_slice)
}
