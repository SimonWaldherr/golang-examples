// Description: Type conversions in Go - numeric, string, and byte-slice conversions
// Tags: type conversion, cast, strconv, string, int, float, byte, rune
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Numeric conversions - explicit cast required in Go
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Printf("int(%d) -> float64(%g) -> uint(%d)\n", i, f, u)

	// Narrowing conversion (information may be lost)
	var big int64 = 1_000_000
	var small int8 = int8(big) // wraps around
	fmt.Printf("int64(%d) -> int8(%d) [overflow wraps]\n", big, small)

	// string <-> []byte (copies the underlying data)
	s := "Hello, 世界"
	b := []byte(s)
	fmt.Printf("string -> []byte: %v\n", b[:5])
	fmt.Printf("[]byte -> string: %q\n", string(b))

	// string <-> []rune (Unicode code points)
	r := []rune(s)
	fmt.Printf("string -> []rune: %v\n", r)
	fmt.Printf("[]rune -> string: %q\n", string(r))
	fmt.Printf("len(string)=%d  len([]rune)=%d\n", len(s), len(r))

	// int <-> string via strconv (NOT via direct cast which gives rune)
	n := 123
	str := strconv.Itoa(n) // int to string
	fmt.Printf("Itoa(%d) = %q\n", n, str)

	parsed, err := strconv.Atoi("456") // string to int
	if err == nil {
		fmt.Printf("Atoi(%q) = %d\n", "456", parsed)
	}

	// FormatFloat / ParseFloat
	fstr := strconv.FormatFloat(3.14159, 'f', 4, 64)
	fmt.Printf("FormatFloat(3.14159) = %q\n", fstr)

	pf, _ := strconv.ParseFloat("2.71828", 64)
	fmt.Printf("ParseFloat(%q) = %g\n", "2.71828", pf)

	// bool conversions
	boolStr := strconv.FormatBool(true)
	fmt.Printf("FormatBool(true) = %q\n", boolStr)

	pb, _ := strconv.ParseBool("false")
	fmt.Printf("ParseBool(%q) = %t\n", "false", pb)
}
