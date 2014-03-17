package main

// import required modules
import (
	"fmt"
)

// a function which add up two integers and returns it as int
func plus(a int, b int) int {
	return a + b
}

// a function which subtract an integer from an integer and returns it as int
func minus(a int, b int) int {
	return a - b
}

// main function
func main() {
	fmt.Println(plus(3, minus(10, 5)))
}
