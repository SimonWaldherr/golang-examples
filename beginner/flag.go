package main

import (
	"flag"
	"fmt"
)

func main() {
	var arg1 int
	var arg2 string
	var arg3 float64
	flag.IntVar(&arg1, "a", 0, "first argument (int)")
	flag.StringVar(&arg2, "b", "lorem", "second argument (string)")
	flag.Float64Var(&arg3, "c", 13.37, "third argument (float)")
	flag.Parse()
	fmt.Printf("flag input was: \n\targ1: %v\n\targ2: %v\n\targ3: %v\n\n", arg1, arg2, arg3)
}
