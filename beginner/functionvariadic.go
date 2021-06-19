package main

import "fmt"

func main() {
	variadic(1, 2, 3, 4, 5)
}

// Variadic functions can have zero or n parameters passed
// The arguments passed to a variadic function are appended to a slice of the same type
func variadic(numbers ...int) {
	fmt.Printf("Type: %T\t Content: %d\n", numbers, numbers)
}
