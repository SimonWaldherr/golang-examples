package main

import (
	"fmt"
)

func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, element := range slice {
		if predicate(element) {
			result = append(result, element)
		}
	}
	return result
}

func main() {
	ints := []int{1, 2, 3, 4, 5}
	evens := Filter(ints, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Even ints:", evens)

	strings := []string{"apple", "banana", "cherry"}
	bWords := Filter(strings, func(s string) bool {
		return len(s) > 5
	})
	fmt.Println("Words longer than 5 characters:", bWords)
}
