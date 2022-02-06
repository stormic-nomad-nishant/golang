/* CONTROL FLOWS

THIS IS THE ERROR HANDLING IN Golang


Defer:
  A defer statement pushes a function call onto a list. The list of saved calls is executed after the surrounding function returns.
  Defer is commonly used to simplify functions that perform various clean-up actions.
  --> DEFER basically invokes a function but delays its execution to some point in future time.

Panic:
  Panic is a built-in function that stops the ordinary flow of control and begins panicking.
  When the function F calls panic, execution of F stops, any deferred functions
  in F are executed normally, and then F returns to its caller.
  --> How a program can panic and can no longer run. How it is invoked


Recover:
  Recover is a built-in function that regains control of a panicking goroutine.
  Recover is only useful inside deferred functions.
  During normal execution, a call to recover will return nil and have no other effect.
  --> How to recover a panicing go program
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
}

/*
Defer executes any functions that are passed into it after the functions
finishes  final statement but before it actually returns.
So the above main function executes as below:

1st: line 28 is printed
2nd: It sees a defer function to call, so it does nothing
3rd: line 30 is printed
4rth: Main method exits
5th: Lastly when go compiler sees that main() method has exited , it checks
if any defer function needs to be called so it goes to line 29 and executes
the defer

--> So defering does not move a function to the end, but it actually moves the
defered function after the (main) function in the above example.
*/
