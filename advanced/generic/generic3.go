package main

import (
	"fmt"
)

func Contains[T comparable](slice []T, item T) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}

func main() {
	ints := []int{1, 2, 3, 4, 5}
	strings := []string{"apple", "banana", "cherry"}

	fmt.Println("Contains 3:", Contains(ints, 3))
	fmt.Println("Contains 6:", Contains(ints, 6))
	fmt.Println("Contains 'banana':", Contains(strings, "banana"))
	fmt.Println("Contains 'grape':", Contains(strings, "grape"))
}
