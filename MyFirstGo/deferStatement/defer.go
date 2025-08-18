package deferStatement

import "fmt"

/*
Defer statements

A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns,
either because the surrounding function executed a return statement, reached the end of its function body, or because the
corresponding goroutine is panicking.

The expression must be a function or method call; it cannot be parenthesized. Calls of built-in functions are
restricted as for expression statements.

Each time a `defer` statement executes, the function value and parameters to the call are evaluated as usual and saved
anew but the actual function is not invoked. Instead, deferred functions are invoked immediately before
*/
func main() {

	//A program is executed line by line and if you add defer statement anywhere then that line won't be executed and it will be executed at the end of the program or method
	//

	defer fmt.Println("World")
	fmt.Println("Hello")

	defer fmt.Println("World")
	fmt.Println("Hello")

	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	
	defer fmt.Println("World")
	fmt.Println("Hello")
	myDefer()
}
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
