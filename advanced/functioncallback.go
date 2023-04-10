// Description: Function callback example
// Tags: function, callback, function callback, function callback, callback, callback function, callback function, function, function callback, function callback, callback, callback function, callback function
package main

import "fmt"

func timestwo(f func(int) int, x int) int {
	return f(x * 2)
}

func main() {
	fmt.Printf("%v\n", timestwo(func(i int) int {
		return i * 2
	}, 32))
}
