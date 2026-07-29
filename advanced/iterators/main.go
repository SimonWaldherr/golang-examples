// Description: Range-over-function iterators and iterator helpers (Go 1.23)
// Tags: iter, iterator, range-over-func, slices, maps, generics, Go 1.23
package main

import (
	"fmt"
	"iter"
	"maps"
	"slices"
)

// Fibonacci returns a reusable pull-free iterator.
func Fibonacci(limit int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for a, b := 0, 1; a <= limit; a, b = b, a+b {
			if !yield(a) {
				return
			}
		}
	}
}

func Filter[T any](input iter.Seq[T], keep func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for value := range input {
			if keep(value) && !yield(value) {
				return
			}
		}
	}
}

func Map[From, To any](input iter.Seq[From], transform func(From) To) iter.Seq[To] {
	return func(yield func(To) bool) {
		for value := range input {
			if !yield(transform(value)) {
				return
			}
		}
	}
}

func main() {
	evenSquares := Map(
		Filter(Fibonacci(100), func(value int) bool { return value%2 == 0 }),
		func(value int) int { return value * value },
	)
	fmt.Println("even Fibonacci squares:", slices.Collect(evenSquares))

	inventory := map[string]int{"gopher": 3, "compiler": 1, "module": 2}
	fmt.Println("sorted map keys:", slices.Sorted(maps.Keys(inventory)))
}
