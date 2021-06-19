package main

import (
	"fmt"
)

// define the function as variable
var patchableFunc = func(i int) int {
	return i * 2
}

func testFunc() {
	// save the original func
	originalFunc := patchableFunc

	// restore the original func via defer at the end of the func
	defer func() {
		patchableFunc = originalFunc
	}()

	// call the function
	fmt.Println(patchableFunc(1))

	// change the function
	patchableFunc = func(i int) int {
		return i * 3
	}

	// call the function again
	fmt.Println(patchableFunc(2))
}

func main() {
	fmt.Println(patchableFunc(3))
	testFunc()
	fmt.Println(patchableFunc(4))
}
