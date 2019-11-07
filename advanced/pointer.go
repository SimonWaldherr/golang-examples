package main

import (
	"fmt"
)

func plusOne(in *int) {
	*in++
}

func main() {
	var value int
	value = 5

	fmt.Printf("value: %v\n", value)
	plusOne(&value)
	fmt.Printf("value: %v\n", value)
	plusOne(&value)
	plusOne(&value)
	plusOne(&value)
	fmt.Printf("value: %v\n", value)
}
