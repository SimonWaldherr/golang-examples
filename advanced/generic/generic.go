// Description: Generic function in Go
// Tags: generic, function, generic function, generic function, function, function generic, function generic
package main

import "fmt"

// Int is a type constraint. It can be used to indicate that a function can accept the
// types listed within the constraint, in this case int, int32 and int64.
type Int interface {
	int | int32 | int64
}

// Float is also a type constraint but describe type float32 and float64.
type Float interface {
	float32 | float64
}

// The Number used the Int and Float constraint. In this case it has the same effect as
//
//	type Number interface {
//		int | int32 | int64 | float32 | float64
//	}
type Number interface {
	Int | Float
}

func main() {
	// Initialize a map for the integer values
	ints := []int64{6, 3, 555}

	// Initialize a map for the float values
	floats := []float64{35.98, 26.99, 933.0001, 3.14}

	fmt.Printf("(Non-Generic) SumInts: %v, SumFloats: %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("(type explicitly specified) SumIntsOrFloats: %v, SumIntsOrFloats: %v\n",
		SumIntsOrFloats[int64](ints),
		SumIntsOrFloats[float64](floats))

	fmt.Printf("(type inferred) SumIntsOrFloats: %v, SumIntsOrFloats: %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	fmt.Printf("(With Constraint) SumIntsGeneric: %v, SumIntsGeneric: %v\n",
		SumIntsGeneric(ints),
		SumFloatsGeneric(floats))

	fmt.Printf("(With Constraint) SumNumbers: %v, SumNumbers: %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}

// SumInts sums the values in s. (Non-Generic)
func SumInts(s []int64) int64 {
	var sum int64
	for _, v := range s {
		sum += v
	}
	return sum
}

// SumFloats sums the values in s. (Non-Generic)
func SumFloats(s []float64) float64 {
	var sum float64
	for _, v := range s {
		sum += v
	}
	return sum
}

// SumIntsGeneric sums the values in s. It supports both Int slice. (Generic)
func SumIntsGeneric[V Int](s []V) V {
	var sum V
	for _, v := range s {
		sum += v
	}
	return sum
}

// SumFloatsGeneric sums the values in s. It supports both Float slice. (Generic)
func SumFloatsGeneric[V Float](s []V) V {
	var sum V
	for _, v := range s {
		sum += v
	}
	return sum
}

// SumIntsOrFloats sums the values in slice s. It supports both int64 and float64 slice. (Generic)
func SumIntsOrFloats[V int64 | float64](s []V) V {
	var sum V
	for _, v := range s {
		sum += v
	}
	return sum
}

// SumNumbers sums the values in s. Its supports any type that fulfill Number constraint. (Generic)
// Note: This function header can also be written as
//
//	func SumNumbers[V Int | Float](s []V) V
func SumNumbers[V Number](s []V) V {
	var sum V
	for _, v := range s {
		sum += v
	}
	return sum
}
