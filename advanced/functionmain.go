// Description: main function and init function
// Tags: main, init, main function, init function, main function, init function, function, function main, function main, main function, main function, function, function init, function init, init function, init function
package main

import (
	"fmt"
)

var i int

func init() {
	fmt.Println("the init function gets started at first")
}

func main() {
RESTART:
	fmt.Println("the main function gets called immediately after")
	if i < 3 {
		i++
		fmt.Println("we can also call the main function ourself (but not init)")
		goto RESTART
	}
}
