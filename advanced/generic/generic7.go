package main

import (
	"fmt"
)

type Pair[T, U any] struct {
	First  T
	Second U
}

func main() {
	intStringPair := Pair[int, string]{First: 1, Second: "one"}
	fmt.Println("Pair:", intStringPair)

	stringFloatPair := Pair[string, float64]{First: "pi", Second: 3.14}
	fmt.Println("Pair:", stringFloatPair)
}
