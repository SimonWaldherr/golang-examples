// Go program to illustrate how a
// function return multiple values
package main

import "fmt"

// myfunc return 3 values of int type
func myfunc(p, q int) (int, int, int) {
	return p - q, p * q, p + q
}

// Main Method
func main() {

	// The return values are assigned into
	// three different variables
	var myvar1, myvar2, myvar3 = myfunc(4, 2)

	// Display the values
	fmt.Printf("Result is: %d", myvar1)
	fmt.Printf("\nResult is: %d", myvar2)
	fmt.Printf("\nResult is: %d", myvar3)
}
