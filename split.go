package main

// import required modules
import (
	"fmt"
	"strings"
)

// declare variables
var variable string
var strarray []string

// main function
func main() {

	// define content of variable
	variable = "Lorem Ipsum Dolor Sit Amet"

	// print variable
	fmt.Println(variable)

	// split variable by space
	strarray = strings.Split(variable, " ")

	// do something as often as the array contains elements
	for i := 0; i < len(strarray); i++ {

		// print the array element defined by i
		fmt.Println(strarray[i])
	}
}
