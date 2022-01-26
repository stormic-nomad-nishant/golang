package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break // BUT THIS WOULD ONLY BREAK Out of INNER LOOP i.e j and not i
				// So we have to use the lable as we show in next
			}
		}
	}
	fmt.Println("with lable:")
Loop:
	for m := 1; m <= 3; m++ {
		for n := 1; n <= 3; n++ {
			fmt.Println(m * n)
			if m*n >= 3 {
				break Loop // So the lable 'Loop' or anything helps u break out of nexted loop
				// So lable tells where we want to break out to?
			}
		}
	}
}
