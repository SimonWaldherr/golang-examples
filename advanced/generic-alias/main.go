// Description: Generic type aliases (Go 1.24)
// Tags: generics, type alias, set, Go 1.24
package main

import (
	"fmt"
	"maps"
	"slices"
)

type Set[Element comparable] map[Element]struct{}

// A generic alias can specialize another generic type without creating a new
// type. NumberSet[int] and Set[int], for example, are identical types.
type NumberSet[Number ~int | ~int64] = Set[Number]

func NewSet[Element comparable](values ...Element) Set[Element] {
	set := make(Set[Element], len(values))
	for _, value := range values {
		set[value] = struct{}{}
	}
	return set
}

func main() {
	var numbers NumberSet[int] = NewSet(5, 3, 5, 1)
	fmt.Println("numbers:", slices.Sorted(maps.Keys(numbers)))

	words := NewSet("go", "gopher", "go")
	fmt.Println("words:", slices.Sorted(maps.Keys(words)))
}
