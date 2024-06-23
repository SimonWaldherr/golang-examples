package main

import (
	"fmt"
)

type Number interface {
	int | float64
}

func Sum[T Number](numbers []T) T {
	var sum T
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	ints := []int{1, 2, 3, 4, 5}
	floats := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

	fmt.Println("Sum of ints:", Sum(ints))
	fmt.Println("Sum of floats:", Sum(floats))
}
