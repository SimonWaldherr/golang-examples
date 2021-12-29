package main

import (
	"fmt"
)

type Number interface {
	int64 | float64
}

func SumInts(m ...int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m ...float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumIntsOrFloats[V int64 | float64](m ...V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbers[V Number](m ...V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	ints := []int64{31, 37, 41, 43, 47, 53, 59}
	floats := []float64{31.17, 37.2, 41.9, 43.002, 47, 53, 59}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints...),
		SumFloats(floats...))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats(ints...),
		SumIntsOrFloats(floats...))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumNumbers(ints...),
		SumNumbers(floats...))
}
