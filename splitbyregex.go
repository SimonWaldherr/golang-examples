package main

import (
	"fmt"
	"regexp"
)

var variable string
var strarray []string
var r = regexp.MustCompile("[^\\s]+")

func main() {
	variable = "Lorem  Ipsum  Dolor   Sit  Amet"
	strarray = r.FindAllString(variable, -1)
	for i := 0; i < len(strarray); i++ {
		fmt.Println(strarray[i])
	}
}
