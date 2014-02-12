package main

import (
	"fmt"
	"os"
)

func main() {
	if os.Args[1] == "Hello" {
		fmt.Println("Hello World")
	} else {
		fmt.Println("Hack the Planet")
	}
}
