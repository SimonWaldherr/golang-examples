// Go program to illustrate how to copy an array by value
package main

import "fmt"

func main() {

	// Creating and initializing an array
	// Using shorthand declaration
	my_arr1 := [5]string{"Scala", "Perl", "Java", "Python", "Go"}

	// Copying the array into new variable
	// Here, the elements are passed by value
	my_arr2 := my_arr1

	fmt.Println("Array_1: ", my_arr1)
	fmt.Println("Array_2:", my_arr2)

	my_arr1[0] = "C++"

	// Here, when we copy an array
	// into another array by value
	// then the changes made in original
	// array do not reflect in the
	// copy of that array
	fmt.Println("\nArray_1: ", my_arr1)
	fmt.Println("Array_2: ", my_arr2)
}
