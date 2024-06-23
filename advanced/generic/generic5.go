package main

import (
	"fmt"
)

func Map[T any, U any](input []T, mapper func(T) U) []U {
	result := make([]U, len(input))
	for i, v := range input {
		result[i] = mapper(v)
	}
	return result
}

func main() {
	ints := []int{1, 2, 3, 4, 5}
	squared := Map(ints, func(n int) int {
		return n * n
	})
	fmt.Println("Squared ints:", squared)

	strings := []string{"a", "b", "c"}
	lengths := Map(strings, func(s string) int {
		return len(s)
	})
	fmt.Println("Lengths of strings:", lengths)
}
