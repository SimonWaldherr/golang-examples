// Fizzbuzz is a common programming test.
// The objective is to print the integers 1 to 100, but:
// 		-for multiples of 3, print fizz instead of the number
// 		-for multiples of 5, print buzz instead of the number
// 		-for multiples of 3 and 5, print fizzbuzz instead of the number

package main

import (
	"fmt"
)

func main() {
	for index := 1; index < 100; index++ {
		switch {
		case index%15 == 0:
			fmt.Println("fizzbuzz")
		case index%3 == 0:
			fmt.Println("fizz")
		case index%5 == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(index)
		}
	}
}
