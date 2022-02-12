package main

import (
	"fmt"
)

func main() {
	fmt.Println("ARRAYS are NOT passed by ref but value : ")
	a := [3]int{1, 2, 3}
	b := a
	fmt.Println(a, b)
	a[1] = 99
	fmt.Println(a, b)
	// The above code will (line 13) will not update b with a[1]=99 because
	// arrays don't have pointer type internally
	fmt.Println("SLICE are also passed by ref - pointers :")
	d := []int{1, 2, 3}
	e := d
	fmt.Println(d, e)
	d[1] = 22
	fmt.Println(d, e)
	// The above code line 21 will have e also having e[1] as 22 because slice internally
	// are passed by references i.e pointers
	println("MAPS are also passed by ref - pointers :")
	k := map[string]string{"foo": "bar", "baz": "buz"}
	l := k
	fmt.Println(k, l)
	k["foo"] = "qos"
	fmt.Println(k, l)
	// SO ALWAYS BE CARE FULL WHEN YOU PASS MAPS AND SLICES in your app
}
