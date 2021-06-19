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
